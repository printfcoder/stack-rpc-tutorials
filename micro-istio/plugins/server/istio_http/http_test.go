package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/mock"
	"github.com/micro/go-micro/server"
)

func TestHTTPServer(t *testing.T) {
	reg := mock.NewRegistry()

	// create server
	srv := NewServer(server.Registry(reg))

	// create server mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`hello world`))
	})

	// create handler
	hd := srv.NewHandler(mux)

	// register handler
	if err := srv.Handle(hd); err != nil {
		t.Fatal(err)
	}

	// start server
	if err := srv.Start(); err != nil {
		t.Fatal(err)
	}

	// register server
	if err := srv.Register(); err != nil {
		t.Fatal(err)
	}

	// lookup server
	service, err := reg.GetService(server.DefaultName)
	if err != nil {
		t.Fatal(err)
	}

	if len(service) != 1 {
		t.Fatalf("Expected 1 service got %d: %+v", len(service), service)
	}

	if len(service[0].Nodes) != 1 {
		t.Fatalf("Expected 1 node got %d: %+v", len(service[0].Nodes), service[0].Nodes)
	}

	// make request
	rsp, err := http.Get(fmt.Sprintf("http://%s:%d", service[0].Nodes[0].Address, service[0].Nodes[0].Port))
	if err != nil {
		t.Fatal(err)
	}
	defer rsp.Body.Close()

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if s := string(b); s != "hello world" {
		t.Fatalf("Expected response %s, got %s", "hello world", s)
	}

	// deregister server
	if err := srv.Deregister(); err != nil {
		t.Fatal(err)
	}

	// try get service
	service, err = reg.GetService(server.DefaultName)
	if err == nil {
		t.Fatalf("Expected %v got %+v", registry.ErrNotFound, service)
	}

	// stop server
	if err := srv.Stop(); err != nil {
		t.Fatal(err)
	}
}
