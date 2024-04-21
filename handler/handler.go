package handler

import (
	"technical-test/domain/responses"
	"technical-test/pkg/token"
	"technical-test/service"

	"github.com/gin-gonic/gin"

	_ "technical-test/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services   *service.Service
	tokenMaker token.Maker
}

func NewHandler(services *service.Service, tokenMaker token.Maker) *Handler {
	return &Handler{
		services:   services,
		tokenMaker: tokenMaker,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h.InitApi(router)

	return router
}

func (h *Handler) InitApi(router *gin.Engine) {
	h.initAuthGroup(router)
	h.initMemberGroup(router)
}

func errorResponse(err error) responses.ErrorMessage {
	return responses.ErrorMessage{
		Message: err.Error(),
	}
}
