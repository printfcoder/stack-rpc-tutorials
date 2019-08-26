// Package http implements a go-micro.Server
package http

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/util/log"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/gorilla/mux"
	"github.com/micro/go-micro/api"
	ha "github.com/micro/go-micro/api/handler/api"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"
)

type httpServer struct {
	api    bool
	server *hServer

	sync.Mutex
	opts     server.Options
	handlers map[string]server.Handler
	exit     chan chan error
}

func init() {
	cmd.DefaultServers["http"] = NewServer
}

func (h *httpServer) handler(service *service, mtype *methodType) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Logf("handle request path:%v", r.URL.Path)

		var argv, replyv reflect.Value

		// Decode the argument value.
		argIsValue := false // if true, need to indirect before calling.
		if mtype.ArgType.Kind() == reflect.Ptr {
			argv = reflect.New(mtype.ArgType.Elem())
		} else {
			argv = reflect.New(mtype.ArgType)
			argIsValue = true
		}

		// get codec
		ct := r.Header.Get("Content-Type")
		if len(ct) == 0 {
			ct = "application/json"
		}
		codec, err := h.newHTTPCodec(ct)
		if err != nil {
			er := errors.InternalServerError("go.micro.server", err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(er.Error()))
			return
		}

		// marshal request
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			er := errors.InternalServerError("go.micro.server", err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(er.Error()))
			return
		}
		r.Body.Close()

		// TODO 支持go-api的api.Request

		// Unmarshal request
		if len(b) > 0 {
			if err := codec.Unmarshal(b, argv.Interface()); err != nil {
				er := errors.InternalServerError("go.micro.server", err.Error())
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(er.Error()))
				return
			}
		}

		if argIsValue {
			argv = argv.Elem()
		}

		// reply value
		replyv = reflect.New(mtype.ReplyType.Elem())

		function := mtype.method.Func
		var returnValues []reflect.Value

		// create a client.Request
		hr := &httpRequest{
			service:     h.opts.Name,
			contentType: ct,
			method:      fmt.Sprintf("%s.%s", service.name, mtype.method.Name),
			request:     argv.Interface(),
		}

		ctx := context.Background()

		// define the handler func
		fn := func(ctx context.Context, req server.Request, rsp interface{}) error {
			returnValues = function.Call([]reflect.Value{service.rcvr, mtype.prepareContext(ctx), reflect.ValueOf(req.Body()), reflect.ValueOf(rsp)})

			// The return value for the method is an error.
			if err := returnValues[0].Interface(); err != nil {
				return err.(error)
			}

			return nil
		}

		// wrap the handler func
		for i := len(h.opts.HdlrWrappers); i > 0; i-- {
			fn = h.opts.HdlrWrappers[i-1](fn)
		}

		// execute the handler
		if appErr := fn(ctx, hr, replyv.Interface()); appErr != nil {
			er := errors.InternalServerError("go.micro.server", appErr.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(er.Error()))
			return
		}

		rsp, err := codec.Marshal(replyv.Interface())
		if err != nil {
			er := errors.InternalServerError("go.micro.server", err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(er.Error()))
			return
		}

		if len(w.Header().Get("Content-Type")) == 0 {
			w.Header().Set("Content-Type", "application/json")
		}

		w.WriteHeader(http.StatusOK)
		w.Write(rsp)
	}, nil
}

func (h *httpServer) apiHandler(service *service, mtype *methodType) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Logf("handle request path:%v", r.URL.Path)
		req, err := requestToProto(r)
		if err != nil {
			er := errors.InternalServerError("go.micro.server", err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			w.Write([]byte(er.Error()))
			return
		}

		var replyv reflect.Value

		// get codec
		ct := r.Header.Get("Content-Type")
		if len(ct) == 0 {
			ct = "application/json"
		}
		codec, err := h.newHTTPCodec(ct)
		if err != nil {
			er := errors.InternalServerError("go.micro.server", err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(er.Error()))
			return
		}

		// reply value
		replyv = reflect.New(mtype.ReplyType.Elem())

		function := mtype.method.Func
		var returnValues []reflect.Value

		// create a client.Request
		hr := &httpRequest{
			service:     h.opts.Name,
			contentType: ct,
			method:      fmt.Sprintf("%s.%s", service.name, mtype.method.Name),
			request:     req,
		}

		ctx := context.Background()

		// define the handler func
		fn := func(ctx context.Context, req server.Request, rsp interface{}) error {
			returnValues = function.Call([]reflect.Value{service.rcvr, mtype.prepareContext(ctx), reflect.ValueOf(req.Body()), reflect.ValueOf(rsp)})

			// The return value for the method is an error.
			if err := returnValues[0].Interface(); err != nil {
				return err.(error)
			}

			return nil
		}

		// wrap the handler func
		for i := len(h.opts.HdlrWrappers); i > 0; i-- {
			fn = h.opts.HdlrWrappers[i-1](fn)
		}

		// execute the handler
		if appErr := fn(ctx, hr, replyv.Interface()); appErr != nil {
			er := errors.InternalServerError("go.micro.server", appErr.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(er.Error()))
			return
		}

		rsp, err := codec.Marshal(replyv.Interface())
		if err != nil {
			er := errors.InternalServerError("go.micro.server", err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(er.Error()))
			return
		}

		if len(w.Header().Get("Content-Type")) == 0 {
			w.Header().Set("Content-Type", "application/json")
		}

		w.WriteHeader(http.StatusOK)
		w.Write(rsp)
	}, nil
}

func (h *httpServer) newHTTPCodec(contentType string) (Codec, error) {
	if c, ok := defaultHTTPCodecs[contentType]; ok {
		return c, nil
	}
	return nil, fmt.Errorf("Unsupported Content-Type: %s", contentType)
}

func (h *httpServer) Options() server.Options {
	h.Lock()
	opts := h.opts
	h.Unlock()
	return opts
}

func (h *httpServer) Init(opts ...server.Option) error {
	h.Lock()
	for _, o := range opts {
		o(&h.opts)
	}
	h.Unlock()
	return nil
}

func (h *httpServer) Handle(handler server.Handler) error {
	h.Lock()
	defer h.Unlock()

	if err := h.server.register(handler.Handler()); err != nil {
		return err
	}

	h.handlers[handler.Name()] = handler

	return nil
}

func (h *httpServer) NewHandler(handler interface{}, opts ...server.HandlerOption) server.Handler {
	return newHttpHandler(handler, opts...)
}

func (h *httpServer) NewSubscriber(topic string, handler interface{}, opts ...server.SubscriberOption) server.Subscriber {
	var options server.SubscriberOptions
	for _, o := range opts {
		o(&options)
	}

	return &httpSubscriber{
		opts:  options,
		topic: topic,
		hd:    handler,
	}
}

func (h *httpServer) Subscribe(s server.Subscriber) error {
	return errors.InternalServerError("go.micro.server", "subscribe is not supported")
}

func (h *httpServer) Register() error {
	return nil
}

func (h *httpServer) Deregister() error {
	return nil
}

func (h *httpServer) Start() error {
	h.Lock()
	opts := h.opts
	serviceMap := h.server.serviceMap
	h.Unlock()

	ln, err := net.Listen("tcp", opts.Address)
	if err != nil {
		return err
	}

	h.Lock()
	h.opts.Address = ln.Addr().String()
	h.Unlock()

	r := mux.NewRouter()
	for _, service := range serviceMap {
		if h.api {
			// micro api
			handler := h.handlers[service.name]
			for _, v := range handler.Options().Metadata {
				if e := api.Decode(v); e == nil {
					continue
				} else {
					if sm := strings.Split(e.Name, "."); len(sm) == 2 && e.Handler == ha.Handler {
						mn := sm[1]
						mt := service.method[mn]
						if mt.stream {
							// TODO stream支持
							continue
						}

						handler, err := h.apiHandler(service, mt)
						if err != nil {
							return err
						}

						for _, path := range e.Path {
							r.HandleFunc(path, handler).Methods(e.Method...)
						}
					}
				}
			}
		} else {
			for mn, mt := range service.method {
				if mt.stream {
					// TODO stream支持
					continue
				}
				handler, err := h.handler(service, mt)
				if err != nil {
					return err
				}

				r.HandleFunc("/"+service.name+"."+mn, handler)
			}
		}
	}

	go http.Serve(ln, r)

	go func() {
		ch := <-h.exit
		ch <- ln.Close()
	}()

	return nil
}

func (h *httpServer) Stop() error {
	ch := make(chan error)
	h.exit <- ch
	return <-ch
}

func (h *httpServer) String() string {
	return "http"
}

func newServer(api bool, opts ...server.Option) server.Server {
	return &httpServer{
		api: api,
		server: &hServer{
			serviceMap: make(map[string]*service),
		},
		opts:     newOptions(opts...),
		exit:     make(chan chan error),
		handlers: make(map[string]server.Handler),
	}
}

func NewServer(opts ...server.Option) server.Server {
	return newServer(false, opts...)
}

func NewApiServer(opts ...server.Option) server.Server {
	return newServer(true, opts...)
}
