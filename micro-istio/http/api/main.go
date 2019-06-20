package main

import (
	"flag"

	httpClient "github.com/micro-in-cn/tutorials/micro-istio/plugins/client/istio_http"
	httpServer "github.com/micro-in-cn/tutorials/micro-istio/plugins/server/istio_http"
	"github.com/micro/cli"
	"github.com/micro/go-micro/api"
	ha "github.com/micro/go-micro/api/handler/api"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"

	apiClient "github.com/micro-in-cn/tutorials/micro-istio/http/api/client"
	"github.com/micro-in-cn/tutorials/micro-istio/http/api/handler"
	example "github.com/micro-in-cn/tutorials/micro-istio/http/api/proto/example"
)

var (
	serverAddr string
	callAddr   string
	cmdHelp    bool
)

func init() {
	flag.StringVar(&serverAddr, "server_address", "0.0.0.0:9080", "server address.")
	flag.StringVar(&callAddr, "client_call_address", ":9080", "client call options address.")
	flag.BoolVar(&cmdHelp, "h", false, "help")
	flag.Parse()
}

func main() {
	if cmdHelp {
		flag.PrintDefaults()
		return
	}

	// TODO 多client需要统一端口，或者在client中hard code
	c := httpClient.NewClient(
		client.ContentType("application/json"),
		func(o *client.Options) {
			o.CallOptions.Address = callAddr
		},
	)
	s := httpServer.NewApiServer(
		server.Address(serverAddr),
	)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.sample"),
		micro.Version("latest"),
		micro.Client(c),
		micro.Server(s),

		// 兼容micro cmd parse
		micro.Flags(cli.StringFlag{
			Name:   "client_call_address",
			EnvVar: "MICRO_CLIENT_CALL_ADDRESS",
			Usage:  " Invalid!!!",
		}),
	)

	service.Options().Cmd.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example),
		api.WithEndpoint(&api.Endpoint{
			// The RPC method
			Name: "Example.Call",
			// The HTTP paths. This can be a POSIX regex
			Path: []string{"/example/call"},
			// The HTTP Methods for this endpoint
			Method: []string{"GET", "POST"},
			// The API handler to use
			Handler: ha.Handler,
		}))

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(apiClient.ExampleWrapper(service)),
	)

	log.Logf("Service Run")

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
