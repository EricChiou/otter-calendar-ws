package middleware

import (
	"errors"
	"otter-calendar/config"
	"otter-calendar/http/jwt"

	"github.com/valyala/fasthttp"
)

const tokenHeader string = "Authorization"
const tokenPrefix string = "Bearer "

func verifyToken(ctx *fasthttp.RequestCtx) (jwt.Payload, error) {
	auth := string(ctx.Request.Header.Peek(tokenHeader))
	if len(auth) < len(tokenPrefix) {
		return jwt.Payload{}, errors.New("token error")
	}

	payload, err := jwt.Verify(auth[len(tokenPrefix):], config.Get().JWTKey)
	if err != nil {
		return payload, errors.New("token error")
	}

	return payload, nil
}
