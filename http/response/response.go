package response

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type Status string

const (
	Success         Status = "success"
	Error           Status = "error"
	FormatError     Status = "formatError"
	TokenError      Status = "tokenError"
	PermissionError Status = "permissionError"
	OperationError  Status = "operationError"
	ServerError     Status = "serverError"
)

type response struct {
	Status  Status      `json:"status"`
	Result  interface{} `json:"result,omitempty"`
	Message string      `json:"message,omitempty"`
	Trace   string      `json:"trace,omitempty"`
}

type record struct {
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
	Total int           `json:"total"`
	List  []interface{} `json:"list"`
}

type send struct{}

func (s send) Success(ctx *fasthttp.RequestCtx, result interface{}) interface{} {
	resp := response{Status: Success, Result: result}
	return s.send(ctx, resp)
}

func (s send) Page(ctx *fasthttp.RequestCtx, page, limit, total int, list []interface{}) interface{} {
	resp := response{Status: Success, Result: record{Page: page, Limit: limit, Total: total, List: list}}
	return s.send(ctx, resp)
}

func (s send) FormatError(ctx *fasthttp.RequestCtx, errorMsg string, err error) interface{} {
	ctx.Response.SetStatusCode(400)
	resp := response{Status: FormatError, Message: errorMsg, Trace: err.Error()}
	return s.send(ctx, resp)
}

func (s send) TokenError(ctx *fasthttp.RequestCtx, errorMsg string, err error) interface{} {
	ctx.Response.SetStatusCode(401)
	resp := response{Status: TokenError, Message: errorMsg, Trace: err.Error()}
	return s.send(ctx, resp)
}

func (s send) PermissionError(ctx *fasthttp.RequestCtx, errorMsg string, err error) interface{} {
	ctx.Response.SetStatusCode(401)
	resp := response{Status: PermissionError, Message: errorMsg, Trace: err.Error()}
	return s.send(ctx, resp)
}

func (s send) OperationError(ctx *fasthttp.RequestCtx, errorMsg string, err error) interface{} {
	ctx.Response.SetStatusCode(403)
	resp := response{Status: OperationError, Message: errorMsg, Trace: err.Error()}
	return s.send(ctx, resp)
}

func (s send) ServerError(ctx *fasthttp.RequestCtx, errorMsg string, err error) interface{} {
	ctx.Response.SetStatusCode(500)
	resp := response{Status: ServerError, Message: errorMsg, Trace: err.Error()}
	return s.send(ctx, resp)
}

func (s send) send(ctx *fasthttp.RequestCtx, resp response) interface{} {
	bytes, _ := json.Marshal(resp)
	fmt.Fprint(ctx, string(bytes))
	return nil
}

var Send = send{}
