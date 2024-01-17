package main

<<<<<<< HEAD
import (
	"log"

	ep "github.com/wrossmorrow/envoy-extproc-sdk-go"
)
=======
import ep "github.com/wrossmorrow/envoy-extproc-sdk-go"
>>>>>>> feat/cmd-args

type noopRequestProcessor struct {
	opts *ep.ProcessingOptions
}

func (s *noopRequestProcessor) GetName() string {
	return "noop"
}

func (s *noopRequestProcessor) GetOptions() *ep.ProcessingOptions {
	return s.opts
}

<<<<<<< HEAD
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
=======
func (s *noopRequestProcessor) ProcessRequestHeaders(ctx *ep.RequestContext, headers map[string][]string) error {
	return ctx.ContinueRequest()
}

func (s *noopRequestProcessor) ProcessRequestBody(ctx *ep.RequestContext, body []byte) error {
	return ctx.ContinueRequest()
}

func (s *noopRequestProcessor) ProcessRequestTrailers(ctx *ep.RequestContext, trailers map[string][]string) error {
	return ctx.ContinueRequest()
}

func (s *noopRequestProcessor) ProcessResponseHeaders(ctx *ep.RequestContext, headers map[string][]string) error {
	return ctx.ContinueRequest()
}

func (s *noopRequestProcessor) ProcessResponseBody(ctx *ep.RequestContext, body []byte) error {
	return ctx.ContinueRequest()
}

func (s *noopRequestProcessor) ProcessResponseTrailers(ctx *ep.RequestContext, trailers map[string][]string) error {
>>>>>>> feat/cmd-args
	return ctx.ContinueRequest()
}

func (s *noopRequestProcessor) Init(opts *ep.ProcessingOptions, nonFlagArgs []string) error {
	s.opts = opts
	return nil
}

func (s *noopRequestProcessor) Finish() {}
