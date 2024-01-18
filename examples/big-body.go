package main

import (
	"log"
	"strconv"

	typev3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	ep "github.com/wrossmorrow/envoy-extproc-sdk-go"
)

type bigBodyRequestProcessor struct {
	opts        *ep.ProcessingOptions
	maxBodySize int64
}

func (s *bigBodyRequestProcessor) GetName() string {
	return "big-body"
}

func (s *bigBodyRequestProcessor) GetOptions() *ep.ProcessingOptions {
	return s.opts
}

const kContentLen = "content-length"

func (s *bigBodyRequestProcessor) ProcessRequestHeaders(ctx *ep.RequestContext, headers map[string][]string, rawValues map[string][]byte) error {
	cancel := func(code int32) error {
		return ctx.CancelRequest(code, map[string]string{}, typev3.StatusCode_name[code])
	}
	raw, ok := rawValues[kContentLen]
	if !ok {
		return cancel(413)
	}

	size, _ := strconv.ParseInt(string(raw), 10, 64)
	if size > s.maxBodySize {
		log.Printf("the body size: %d exceeded the maximum size: %d\n", size, s.maxBodySize)
		return cancel(413)
	}

	return ctx.ContinueRequest()
}

func (s *bigBodyRequestProcessor) ProcessRequestBody(ctx *ep.RequestContext, body []byte) error {
	return ctx.ContinueRequest()
}

func (s *bigBodyRequestProcessor) ProcessRequestTrailers(ctx *ep.RequestContext, trailers map[string][]string, rawValues map[string][]byte) error {
	return ctx.ContinueRequest()
}

func (s *bigBodyRequestProcessor) ProcessResponseHeaders(ctx *ep.RequestContext, headers map[string][]string, rawValues map[string][]byte) error {
	return ctx.ContinueRequest()
}

func (s *bigBodyRequestProcessor) ProcessResponseBody(ctx *ep.RequestContext, body []byte) error {
	return ctx.ContinueRequest()
}

func (s *bigBodyRequestProcessor) ProcessResponseTrailers(ctx *ep.RequestContext, trailers map[string][]string, rawValues map[string][]byte) error {
	return ctx.ContinueRequest()
}

const kBodySize = "body-size"

func (s *bigBodyRequestProcessor) Init(opts *ep.ProcessingOptions, nonFlagArgs []string) error {
	s.opts = opts
	s.maxBodySize = 16

	var (
		i           int
		err         error
		maxBodySize int64
	)

	nArgs := len(nonFlagArgs)
	for ; i < nArgs-1; i++ {
		if nonFlagArgs[i] == kBodySize {
			break
		}
	}

	if i == nArgs {
		log.Printf("the argument: 'body-size' is missing, use the default.\n")
		return nil
	}

	maxBodySize, err = strconv.ParseInt(nonFlagArgs[i+1], 10, 64)
	if err != nil {
		log.Printf("parse the value for parameter: 'body-size' is failed: %v,use the default.\n", err.Error())
		return nil
	}

	if maxBodySize > 0 {
		s.maxBodySize = maxBodySize
		log.Printf("the max body size is: %d.\n", s.maxBodySize)
	}

	return nil
}

func (s *bigBodyRequestProcessor) Finish() {}
