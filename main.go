package main

import (
	"technical-test/handler"
	"technical-test/mapper"
	"technical-test/pkg/database/migration"
	"technical-test/pkg/database/postgres"
	"technical-test/pkg/dotenv"
	"technical-test/pkg/hashing"
	"technical-test/pkg/logger"
	"technical-test/pkg/token"
	"technical-test/repository"
	"technical-test/service"

	"go.uber.org/zap"
)

// @title K-Link Technical Test API
// @version 1.0
// @description This is K-Link Technical Test API for Developer position
// @termsOfService http://swagger.io/terms/

// @contact.name Alif Dewantara
// @contact.url http://github.com/alifdwt
// @contact.email aputradewantara@gmail.com

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
func main() {
	log, err := logger.NewLogger()
	if err != nil {
		log.Error("Error while initializing logger", zap.Error(err))
	}

	config, err := dotenv.LoadConfig(".")
	if err != nil {
		log.Error("Error while load environtment variables", zap.Error(err))
	}

	db, err := postgres.NewClient(
		config.DB_HOST,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_NAME,
		config.DB_PORT,
		config.APP_TIMEZONE,
	)
	if err != nil {
		log.Error("Error while connecting to database", zap.Error(err))
	}

	err = migration.RunMigration(db)
	if err != nil {
		log.Error("Error while running migration", zap.Error(err))
	}

	hashing := hashing.NewHashingPassword()
	repository := repository.NewRepository(db, config)

	tokenMaker, err := token.NewJWTMaker(config.TOKEN_SYMETRIC_KEY)
	if err != nil {
		log.Error("Error while creating token maker", zap.Error(err))
	}

	mapper := mapper.NewMapper()

	service := service.NewService(service.Deps{
		Config:     config,
		Repository: repository,
		Hashing:    *hashing,
		TokenMaker: tokenMaker,
		Logger:     *log,
		Mapper:     *mapper,
	})

	myHandler := handler.NewHandler(service, tokenMaker)

	err = myHandler.Init().Run(config.SERVER_ADDRESS)
	if err != nil {
		log.Error("Cannot start server", zap.Error(err))
	}
}
