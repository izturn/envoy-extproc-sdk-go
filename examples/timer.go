package main

import (
	"log"
	"strconv"
	"time"

	ep "github.com/wrossmorrow/envoy-extproc-sdk-go"
)

type timerRequestProcessor struct {
	opts *ep.ProcessingOptions
}

func (s *timerRequestProcessor) GetName() string {
	return "timer"
}

func (s *timerRequestProcessor) GetOptions() *ep.ProcessingOptions {
	return s.opts
}

<<<<<<< HEAD
func (s timerRequestProcessor) ProcessRequestHeaders(ctx *ep.RequestContext, headers map[string][]string) error {
	log.Println("timer ProcessRequestHeaders")
=======
func (s *timerRequestProcessor) ProcessRequestHeaders(ctx *ep.RequestContext, headers map[string][]string) error {
>>>>>>> feat/cmd-args
	ctx.OverwriteHeader("x-extproc-started-ns", strconv.FormatInt(ctx.Started.UnixNano(), 10))
	return ctx.ContinueRequest()
}

<<<<<<< HEAD
func (s timerRequestProcessor) ProcessRequestBody(ctx *ep.RequestContext, body []byte) error {
	log.Println("timer ProcessRequestBody")
	return ctx.ContinueRequest()
}

func (s timerRequestProcessor) ProcessRequestTrailers(ctx *ep.RequestContext, trailers map[string][]string) error {
	log.Println("timer ProcessRequestTrailers")
	return ctx.ContinueRequest()
}

func (s timerRequestProcessor) ProcessResponseHeaders(ctx *ep.RequestContext, headers map[string][]string) error {
	log.Println("timer ProcessResponseHeaders")
=======
func (s *timerRequestProcessor) ProcessRequestBody(ctx *ep.RequestContext, body []byte) error {
	return ctx.ContinueRequest()
}

func (s *timerRequestProcessor) ProcessRequestTrailers(ctx *ep.RequestContext, trailers map[string][]string) error {
	return ctx.ContinueRequest()
}

func (s *timerRequestProcessor) ProcessResponseHeaders(ctx *ep.RequestContext, headers map[string][]string) error {
>>>>>>> feat/cmd-args
	finished := time.Now()
	duration := time.Since(ctx.Started)

	ctx.AddHeader("x-extproc-started-ns", strconv.FormatInt(ctx.Started.UnixNano(), 10))
	ctx.AddHeader("x-extproc-finished-ns", strconv.FormatInt(finished.UnixNano(), 10))
	ctx.AddHeader("x-upstream-duration-ns", strconv.FormatInt(duration.Nanoseconds(), 10))

	return ctx.ContinueRequest()
}

<<<<<<< HEAD
func (s timerRequestProcessor) ProcessResponseBody(ctx *ep.RequestContext, body []byte) error {
	log.Println("timer ProcessResponseBody")
=======
func (s *timerRequestProcessor) ProcessResponseBody(ctx *ep.RequestContext, body []byte) error {
>>>>>>> feat/cmd-args
	finished := time.Now()
	duration := time.Since(ctx.Started)

	ctx.OverwriteHeader("x-extproc-started-ns", strconv.FormatInt(ctx.Started.UnixNano(), 10))
	ctx.OverwriteHeader("x-extproc-finished-ns", strconv.FormatInt(finished.UnixNano(), 10))
	ctx.OverwriteHeader("x-upstream-duration-ns", strconv.FormatInt(duration.Nanoseconds(), 10))

	return ctx.ContinueRequest()
}

<<<<<<< HEAD
func (s timerRequestProcessor) ProcessResponseTrailers(ctx *ep.RequestContext, trailers map[string][]string) error {
	log.Println("timer ProcessResponseTrailers")
=======
func (s *timerRequestProcessor) ProcessResponseTrailers(ctx *ep.RequestContext, trailers map[string][]string) error {
>>>>>>> feat/cmd-args
	return ctx.ContinueRequest()
}

func (s *timerRequestProcessor) Init(opts *ep.ProcessingOptions, nonFlagArgs []string) error {
	s.opts = opts
	return nil
}

func (s *timerRequestProcessor) Finish() {}
