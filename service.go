package busybee

type Service interface {
	Name() string
	Description() string
	DecodeRequest(Decoder) (interface{}, error)
	DecodeResponse(Decoder) (interface{}, error)
	BuildEndpoint() Endpoint
}
