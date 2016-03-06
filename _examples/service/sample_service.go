package service

import (
	"fmt"

	bb "github.com/kazukgw/busybee"
	"golang.org/x/net/context"
)

type SampleServiceRequest struct {
	Arg1 string
	Arg2 string
}

type SampleServiceResponse struct {
	Ret1 string
	Ret2 string
}

type SampleService struct{}

func (svc SampleService) Name() string {
	return "SampleService"
}

func (svc SampleService) Description() string {
	return "This Service is just only sample."
}

func (svc SampleService) DecodeRequest(decoder bb.Decoder) (interface{}, error) {
	req := SampleServiceRequest{}
	err := decoder.Decode(&req)
	return req, err
}

func (svc SampleService) DecodeResponse(decoder bb.Decoder) (interface{}, error) {
	res := SampleServiceResponse{}
	err := decoder.Decode(&res)
	return res, err
}

func (svc SampleService) BuildEndpoint() bb.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SampleServiceRequest)
		fmt.Println(ctx.Value("Content-type"))
		fmt.Println("arg1:", req.Arg1)
		fmt.Println("arg2:", req.Arg2)
		return SampleServiceResponse{Ret1: "hoge", Ret2: "fuga"}, nil
	}
}
