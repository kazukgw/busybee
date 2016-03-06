package busybee

import (
	"io"
)

type Codec interface {
	NewDecoder(io.Reader) Decoder
	NewEncoder(io.Writer) Encoder
	Unmarshal([]byte, interface{}) error
	Marshal(interface{}) ([]byte, error)
}

type Decoder interface {
	Decode(interface{}) error
}

type Encoder interface {
	Encode(interface{}) error
}
