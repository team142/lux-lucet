package lulu

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

//StartRestServer starts a web server that listens on addr and responds with a json body of the server state
func StartRestServer(addr string, healthServer *HealthServer) {
	err := fasthttp.ListenAndServe(addr, func(ctx *fasthttp.RequestCtx) {
		handleRequest(healthServer, ctx)
	})
	if err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
func handleRequest(healthServer *HealthServer, ctx *fasthttp.RequestCtx) {
	health := healthServer.Query()
	b, err := json.Marshal(health)
	if err != nil {
		log.Println(err.Error())
		respondToReq(ctx, 500, "text/plain; charset=utf8", "Internal server error!\n\n")
		return
	}
	respondToReq(ctx, 200, "application/json; charset=utf8", string(b))
	return

}

func respondToReq(ctx *fasthttp.RequestCtx, code int, contentType, body string) {
	ctx.Response.SetStatusCode(code)
	ctx.SetContentType(contentType)
	_, err := fmt.Fprintf(ctx, body)
	if err != nil {
		log.Println(err.Error())
	}

}
