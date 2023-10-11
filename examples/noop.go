package main

import (
	"log"

	ep "github.com/wrossmorrow/envoy-extproc-sdk-go"
)

type noopRequestProcessor struct{}

func (s noopRequestProcessor) GetName() string {
	return "noop"
}

func (s noopRequestProcessor) GetOptions() *ep.ProcessingOptions {
	opts := ep.NewOptions()
	opts.LogStream = true
	opts.LogPhases = true
	opts.UpdateExtProcHeader = true
	opts.UpdateDurationHeader = true
	return opts
}

func (s noopRequestProcessor) ProcessRequestHeaders(ctx *ep.RequestContext, headers map[string][]string) error {
	log.Println("noop ProcessRequestHeaders")
	return ctx.ContinueRequest()
}

func (s noopRequestProcessor) ProcessRequestBody(ctx *ep.RequestContext, body []byte) error {
	log.Println("noop ProcessRequestBody")
	return ctx.ContinueRequest()
}

func (s noopRequestProcessor) ProcessRequestTrailers(ctx *ep.RequestContext, trailers map[string][]string) error {
	log.Println("noop ProcessRequestTrailers")
	return ctx.ContinueRequest()
}

func (s noopRequestProcessor) ProcessResponseHeaders(ctx *ep.RequestContext, headers map[string][]string) error {
	log.Println("noop ProcessResponseHeaders")
	return ctx.ContinueRequest()
}

func (s noopRequestProcessor) ProcessResponseBody(ctx *ep.RequestContext, body []byte) error {
	log.Println("noop ProcessResponseBody")
	return ctx.ContinueRequest()
}

func (s noopRequestProcessor) ProcessResponseTrailers(ctx *ep.RequestContext, trailers map[string][]string) error {
	log.Println("noop ProcessResponseTrailers")
	return ctx.ContinueRequest()
}
