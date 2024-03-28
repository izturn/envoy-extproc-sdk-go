package main

import (
	"regexp"
	"strings"

	ep "github.com/wrossmorrow/envoy-extproc-sdk-go"
	extproc "github.com/wrossmorrow/envoy-extproc-sdk-go"
)

type echoRequestProcessor struct {
	opts *ep.ProcessingOptions
}

func joinHeaders(mvhs map[string][]string) map[string]extproc.HeaderValue {
	hs := make(map[string]extproc.HeaderValue)
	for n, vs := range mvhs {
		hs[n] = extproc.HeaderValue{Value: strings.Join(vs, ",")}
	}
	return hs
}

func (s *echoRequestProcessor) GetName() string {
	return "echo"
}

func (s *echoRequestProcessor) GetOptions() *ep.ProcessingOptions {
	return s.opts
}

func (s *echoRequestProcessor) PreprocessContext(ctx *ep.RequestContext) error {
	echoPathRx, _ := regexp.Compile("/echo/.*")
	ctx.SetValue("echoPath", echoPathRx)
	return nil
}

func (s *echoRequestProcessor) ProcessRequestHeaders(ctx *ep.RequestContext, headers ep.AllHeaders) error {
	match, _ := regexp.MatchString("/echo/.*", ctx.Path)
	if !match {
		return ctx.ContinueRequest()
	}

	if ctx.EndOfStream {
		return ctx.CancelRequest(200, joinHeaders(ctx.AllHeaders.Headers), "")
	}
	return ctx.ContinueRequest()

	// switch ctx.Method {
	// // cancel request when there is no body
	// case "HEAD", "OPTIONS", "GET", "DELETE":
	// 	return ctx.CancelRequest(200, joinHeaders(ctx.Headers), "")
	// default:
	// 	break
	// }
	// return ctx.ContinueRequest()
}

func (s *echoRequestProcessor) ProcessRequestBody(ctx *ep.RequestContext, body []byte) error {
	match, _ := regexp.MatchString("/echo/.*", ctx.Path)
	if !match {
		return ctx.ContinueRequest()
	}
	return ctx.CancelRequest(200, joinHeaders(ctx.AllHeaders.Headers), string(body))
}

func (s *echoRequestProcessor) ProcessRequestTrailers(ctx *ep.RequestContext, trailers ep.AllHeaders) error {
	return ctx.ContinueRequest()
}

func (s *echoRequestProcessor) ProcessResponseHeaders(ctx *ep.RequestContext, headers ep.AllHeaders) error {
	return ctx.ContinueRequest()
}

func (s *echoRequestProcessor) ProcessResponseBody(ctx *ep.RequestContext, body []byte) error {
	return ctx.ContinueRequest()
}

func (s *echoRequestProcessor) ProcessResponseTrailers(ctx *ep.RequestContext, trailers ep.AllHeaders) error {
	return ctx.ContinueRequest()
}

func (s *echoRequestProcessor) Init(opts *ep.ProcessingOptions, nonFlagArgs []string) error {
	s.opts = opts
	return nil
}

func (s *echoRequestProcessor) Finish() {}
