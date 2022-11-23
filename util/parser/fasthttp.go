package parser

import "github.com/valyala/fasthttp"

func ParseQueryParam(ctx *fasthttp.RequestCtx) map[string][]string {
	params := map[string][]string{}
	ctx.QueryArgs().VisitAll(func(key []byte, value []byte) {
		params[string(key)] = []string{string(value)}
	})
	return params
}
