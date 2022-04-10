package router

import (
	"otter-calendar/http/middleware"

	"github.com/EricChiou/httprouter"
	"github.com/valyala/fasthttp"
)

func Init() {
	initUserAPI()
	initEventAPI()
}

func ListenAndServe(serverName, port string) error {
	return newFHServer(serverName).ListenAndServe(":" + port)
}

func ListenAndServeTLS(serverName, port, certPath, keyPath string) error {
	return newFHServer(serverName).ListenAndServeTLS(":"+port, certPath, keyPath)
}

func SetHeader(key string, value string) {
	httprouter.SetHeader(key, value)
}

func newFHServer(serverName string) *fasthttp.Server {
	return &fasthttp.Server{
		Name:    serverName,
		Handler: httprouter.FasthttpHandler(),
	}
}

func get(path string, needToken bool, run func(middleware.WebInput)) {
	httprouter.Get(path, func(ctx *httprouter.Context) {
		middleware.Set(ctx, needToken, run)
	})
}

func post(path string, needToken bool, run func(middleware.WebInput)) {
	httprouter.Post(path, func(ctx *httprouter.Context) {
		middleware.Set(ctx, needToken, run)
	})
}

func put(path string, needToken bool, run func(middleware.WebInput)) {
	httprouter.Put(path, func(ctx *httprouter.Context) {
		middleware.Set(ctx, needToken, run)
	})
}

func delete(path string, needToken bool, run func(middleware.WebInput)) {
	httprouter.Delete(path, func(ctx *httprouter.Context) {
		middleware.Set(ctx, needToken, run)
	})
}

func patch(path string, needToken bool, run func(middleware.WebInput)) {
	httprouter.Patch(path, func(ctx *httprouter.Context) {
		middleware.Set(ctx, needToken, run)
	})
}

func head(path string, needToken bool, run func(middleware.WebInput)) {
	httprouter.Head(path, func(ctx *httprouter.Context) {
		middleware.Set(ctx, needToken, run)
	})
}

func options(path string, needToken bool, run func(middleware.WebInput)) {
	httprouter.Options(path, func(ctx *httprouter.Context) {
		middleware.Set(ctx, needToken, run)
	})
}
