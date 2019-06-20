package http

import (
	"reflect"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
)

type httpHandler struct {
	name      string
	opts      server.HandlerOptions
	endpoints []*registry.Endpoint
	handler   interface{}
}

func newHttpHandler(handler interface{}, opts ...server.HandlerOption) server.Handler {

	options := server.HandlerOptions{
		Metadata: make(map[string]map[string]string),
	}

	for _, o := range opts {
		o(&options)
	}

	typ := reflect.TypeOf(handler)
	hdlr := reflect.ValueOf(handler)
	name := reflect.Indirect(hdlr).Type().Name()

	var endpoints []*registry.Endpoint

	for m := 0; m < typ.NumMethod(); m++ {
		if e := extractEndpoint(typ.Method(m)); e != nil {
			e.Name = name + "." + e.Name

			for k, v := range options.Metadata[e.Name] {
				e.Metadata[k] = v
			}

			endpoints = append(endpoints, e)
		}
	}

	return &httpHandler{
		name:      name,
		handler:   handler,
		endpoints: endpoints,
		opts:      options,
	}
}

func (h *httpHandler) Name() string {
	return h.name
}

func (h *httpHandler) Handler() interface{} {
	return h.handler
}

func (h *httpHandler) Endpoints() []*registry.Endpoint {
	return h.endpoints
}

func (h *httpHandler) Options() server.HandlerOptions {
	return h.opts
}
