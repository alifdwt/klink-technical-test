package token

import "time"

type Maker interface {
	CreateToken(userId int, username string, privilege string, audience string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
