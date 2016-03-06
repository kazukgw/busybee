package busybee

import (
	"encoding/json"
	"io"
)

type CodecJSON struct{}

func (codec CodecJSON) NewDecoder(r io.Reader) Decoder {
	return json.NewDecoder(r)
}

func (codec CodecJSON) NewEncoder(w io.Writer) Encoder {
	return json.NewEncoder(w)
}

func (codec CodecJSON) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (codec CodecJSON) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
