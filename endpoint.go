package busybee

import (
	"golang.org/x/net/context"
)

type Endpoint func(ctx context.Context, request interface{}) (interface{}, error)

type Middleware func(Endpoint) Endpoint
