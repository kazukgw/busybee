package main

import (
	"fmt"

	bb "github.com/kazukgw/busybee"
	"golang.org/x/net/context"
)

func SampleMw(ep bb.Endpoint) bb.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {
		fmt.Println("in sample mw: before")
		res, err := ep(ctx, request)
		fmt.Println("in sample mw: after")
		return res, err
	}
}
