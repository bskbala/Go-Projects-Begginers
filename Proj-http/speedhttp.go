package main

import (
	"github.com/valyala/fasthttp"
)

//FastHTTPHandler  creates fasthttp.RequestHandler.
//Routes
//GET/ping return "Pong"

func FastHTTPHandler() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/ping":
			ctx.SetStatusCode(fasthttp.StatusOK)
			ctx.SetContentType("text/plain")
			ctx.SetBody([]byte("pong"))
		default:
			panic("unsupported Path")

		}

	}

}
