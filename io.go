package molu

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
		fmt.Fprintf(ctx, "Internal server error!\n\n")
		ctx.SetContentType("text/plain; charset=utf8")
		ctx.Response.SetStatusCode(500)
		return
	}
	fmt.Fprint(ctx, string(b))
	ctx.SetContentType("application/json; charset=utf8")
	ctx.Response.SetStatusCode(200)
	return

}
