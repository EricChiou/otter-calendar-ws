package middleware

import (
	"otter-calendar/http/jwt"
	"otter-calendar/http/response"

	"github.com/EricChiou/httprouter"
	"github.com/valyala/fasthttp"
)

type WebInput struct {
	Ctx     *fasthttp.RequestCtx
	Context *httprouter.Context
	Payload jwt.Payload
}

func Set(context *httprouter.Context, needToken bool, run func(WebInput)) {
	webInput := WebInput{
		Ctx:     context.Ctx,
		Context: context,
	}

	// check token
	payload, err := verifyToken(context.Ctx)
	webInput.Payload = payload
	if needToken && err != nil {
		response.Send.TokenError(context.Ctx, "", err)
		return
	}

	run(webInput)
}
