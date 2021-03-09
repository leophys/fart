package main

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			req, err := httputil.DumpRequest(r, true)
			if err != nil {
				ctx.Warnf("Error on request dumping: %s\n", err)
				return r, nil
			}
			ctx.Logf("Request:\n---------\n%s\n---------\n", string(req))
			return r, nil
		})
	proxy.OnResponse().DoFunc(
		func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			resp, err := httputil.DumpResponse(r, true)
			if err != nil {
				ctx.Warnf("Error on response dumping: %s\n", err)
				return r
			}
			ctx.Logf("Response:\n---------\n%s\n---------\n", string(resp))
			return r
		})
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
