package main

import (
	"fmt"

	bb "github.com/kazukgw/busybee"
	"github.com/kazukgw/busybee/_examples/service"
)

func main() {
	svc := service.SampleService{}
	discov := bb.NewStaticServiceDiscovery()
	discov.AddService("http://localhost:8080/"+svc.Name(), svc)
	headerValues := map[string]string{
		"Content-type": "application/json",
	}

	c := bb.NewClient(svc, discov, bb.CodecJSON{}, headerValues)
	req := service.SampleServiceRequest{}
	req.Arg1 = "hoge"
	req.Arg2 = "fuga"
	c.DoWithHandler(req, func(result bb.Result) {
		if result.Error != nil {
			panic(result.Error.Error())
		}
		res := result.Response.(service.SampleServiceResponse)
		fmt.Printf("%#v\n", &res)
	})

	for {
		select {}
	}
}
