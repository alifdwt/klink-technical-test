package token

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	Username  string `json:"username"`
	Privilege string `json:"privilege"`
	jwt.RegisteredClaims
}

var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token is invalid")
)

func NewPayload(userId int, username string, privilege string, audience string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		username,
		privilege,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    strconv.Itoa(int(time.Now().Unix())),
			Subject:   strconv.Itoa(int(userId)),
			ID:        tokenId.String(),
		},
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt.Time) {
		return ErrExpiredToken
	}
	return nil
}
