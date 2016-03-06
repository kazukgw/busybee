package busybee

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

type Client struct {
	Service
	ServiceDiscovery
	Codec
	HeaderValues map[string]string
}

type Result struct {
	Error    error
	Response interface{}
}

func NewClient(
	svc Service,
	discov ServiceDiscovery,
	codec Codec,
	headerValues map[string]string,
) Client {
	return Client{svc, discov, codec, headerValues}
}

func (c Client) Do(request interface{}) <-chan Result {
	resultChan := make(chan Result, 1)
	go func() {
		url := c.ServiceDiscovery.Discover(c.Service)
		data, err := c.Codec.Marshal(request)
		if err != nil {
			resultChan <- Result{Error: CodecError{err}}
			return
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			resultChan <- Result{Error: err}
			return
		}
		for k, v := range c.HeaderValues {
			req.Header.Set(k, v)
		}

		cli := http.Client{}
		res, err := cli.Do(req)
		if err != nil {
			resultChan <- Result{Error: ClientError{err}}
			return
		}

		byteBuf := []byte{}
		buf := bytes.NewBuffer(byteBuf)
		io.Copy(buf, res.Body)
		res.Body.Close()
		reader := bytes.NewReader(buf.Bytes())
		decoder := c.Codec.NewDecoder(reader)

		iserr, errRes := isError(decoder)
		if iserr {
			resultChan <- Result{Error: ServiceError{errors.New(errRes.Error)}}
			return
		}

		reader.Seek(0, 0)
		response, err := c.Service.DecodeResponse(decoder)
		if err != nil {
			resultChan <- Result{Error: CodecError{err}}
			return
		}
		resultChan <- Result{Response: response}
	}()
	return resultChan
}

func (c Client) DoWithHandler(request interface{}, handler func(result Result)) {
	resultChan := c.Do(request)
	go func() {
		result := <-resultChan
		handler(result)
	}()
}

type ErrorResponse struct {
	Error string
}

func isError(decoder Decoder) (bool, ErrorResponse) {
	errres := ErrorResponse{}
	decoder.Decode(&errres)
	return errres != ErrorResponse{}, errres
}

type ServiceError struct {
	Err error
}

func (se ServiceError) Error() string {
	return se.Err.Error()
}

type CodecError struct {
	Err error
}

func (ce CodecError) Error() string {
	return ce.Err.Error()
}

type ClientError struct {
	Err error
}

func (ce ClientError) Error() string {
	return ce.Err.Error()
}
