package server

import (
	"net/http"
	"net/http/httputil"

	"github.com/elazarl/goproxy"
	fmt "gopkg.in/ffmt.v1"
)

type ServeParams struct {
	BindAddr        string
	CtrlAddr        string
	WebsocketAddr   string
	BlacklistTarget []string
	WhitelistTarget []string
}

type Session struct {
	intercepted bool
}

func Server(params ServeParams) error {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			return onRequest(&params, r, ctx)
		})
	proxy.OnResponse().DoFunc(
		func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			return onResponse(r, ctx)
		})

	return http.ListenAndServe(params.BindAddr, proxy)
}

func onRequest(params *ServeParams, r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	var s Session
	var req []byte
	var dumpErr error

	fmt.Pjson(params)
	fmt.Print(r.Host)

	if 0 == len(params.WhitelistTarget) {
		for _, p := range params.BlacklistTarget {
			if p == r.Host {
				return r, nil
			}
		}

		s.intercepted = true

	} else {
		for _, p := range params.WhitelistTarget {
			if p == r.Host {
				s.intercepted = true
			}
		}
	}

	if s.intercepted {
		req, dumpErr = httputil.DumpRequest(r, true)
		if dumpErr != nil {
			ctx.Warnf("Error on request dumping: %s\n", dumpErr)
			return r, nil
		}

		ctx.UserData = s
		ctx.Logf("Request:\n---------\n%s\n---------\n", string(req))
	}

	return r, nil
}

func onResponse(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	resp, err := httputil.DumpResponse(r, true)
	if err != nil {
		ctx.Warnf("Error on response dumping: %s\n", err)

		return r
	}

	switch ctx.UserData.(type) {
	case Session:
		ctx.Logf("Response:\n---------\n%s\n---------\n", string(resp))

	default:
	}

	return r
}
