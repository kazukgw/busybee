package main

import (
	"net/http"

	bb "github.com/kazukgw/busybee"
	"github.com/kazukgw/busybee/_examples/service"
	"golang.org/x/net/context"
)

func main() {
	mux := http.DefaultServeMux
	ctx := context.Background()
	codec := bb.CodecJSON{}

	s := bb.NewServer(mux, ctx, codec)
	s.Context = context.Background()
	s.AddService(service.SampleService{})
	s.Use(SampleMw)
	s.AddContextValueFromHeader("Content-type")
	s.Build()

	http.ListenAndServe(":8080", nil)
}
