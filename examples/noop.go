package main

import (
	ep "github.com/wrossmorrow/envoy-extproc-sdk-go"
	"log"
	"sync"
)

type noopRequestProcessor struct {
	opts       *ep.ProcessingOptions
	callCounts map[string]int
	mu         sync.Mutex
}

func (s *noopRequestProcessor) GetName() string {
	return "noop"
}

func (s *noopRequestProcessor) GetOptions() *ep.ProcessingOptions {
	return s.opts
}

func (s *noopRequestProcessor) incrementAndLog(methodName string) {
	s.mu.Lock()
	s.callCounts[methodName]++
	log.Printf("Method %s called %d times\n", methodName, s.callCounts[methodName])
	s.mu.Unlock()
}

func (s *noopRequestProcessor) ProcessRequestHeaders(ctx *ep.RequestContext, headers map[string][]string, headerRawValues map[string][]byte) error {
	s.incrementAndLog("ProcessRequestHeaders")
	return ctx.ContinueRequest()
}

func (s *noopRequestProcessor) ProcessRequestBody(ctx *ep.RequestContext, body []byte) error {
	s.incrementAndLog("ProcessRequestBody")
	return ctx.ContinueRequest()
}

func (s *noopRequestProcessor) ProcessRequestTrailers(ctx *ep.RequestContext, trailers map[string][]string, rawValues map[string][]byte) error {
	s.incrementAndLog("ProcessRequestTrailers")
	return ctx.ContinueRequest()
}

func (s *noopRequestProcessor) ProcessResponseHeaders(ctx *ep.RequestContext, headers map[string][]string, rawValues map[string][]byte) error {
	s.incrementAndLog("ProcessResponseHeaders")
	return ctx.ContinueRequest()
}

func (s *noopRequestProcessor) ProcessResponseBody(ctx *ep.RequestContext, body []byte) error {
	s.incrementAndLog("ProcessResponseBody")
	return ctx.ContinueRequest()
}

func (s *noopRequestProcessor) ProcessResponseTrailers(ctx *ep.RequestContext, trailers map[string][]string, rawValues map[string][]byte) error {
	s.incrementAndLog("ProcessResponseTrailers")
	return ctx.ContinueRequest()
}

func (s *noopRequestProcessor) Init(opts *ep.ProcessingOptions, nonFlagArgs []string) error {
	s.opts = opts
	s.callCounts = make(map[string]int)
	return nil
}

func (s *noopRequestProcessor) Finish() {}
