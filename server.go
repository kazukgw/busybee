package busybee

import (
	"net/http"

	"golang.org/x/net/context"
)

func NewServer(mux *http.ServeMux, ctx context.Context, codec Codec) Server {
	s := Server{}
	s.Mux = mux
	s.Context = ctx
	s.Codec = codec
	return s
}

type Server struct {
	Mux *http.ServeMux
	context.Context
	Codec
	Services      []Service
	Middlewares   []Middleware
	ContextValues []string
}

func (s *Server) AddService(svc Service) {
	s.Services = append(s.Services, svc)
}

func (s *Server) Use(mw Middleware) {
	s.Middlewares = append(s.Middlewares, mw)
}

func (s *Server) AddContextValueFromHeader(key string) {
	s.ContextValues = append(s.ContextValues, key)
}

func (s *Server) Build() {
	for _, svc := range s.Services {
		ep := s.BuildEndpoint(svc)
		s.AddRoute(svc, s.BuildHandlerFunc(svc, ep))
	}
}

func (s *Server) ContextFromHeader(
	ctx context.Context,
	r *http.Request,
) context.Context {
	h := r.Header
	c := ctx
	for _, key := range s.ContextValues {
		c = context.WithValue(c, key, h.Get(key))
	}
	return c
}

func (s *Server) BuildEndpoint(svc Service) Endpoint {
	next := svc.BuildEndpoint()
	for i := len(s.Middlewares) - 1; i >= 0; i-- {
		next = s.Middlewares[i](next)
	}
	return next
}

func (s *Server) BuildHandlerFunc(
	svc Service,
	ep Endpoint,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithCancel(s.Context)
		ctx = s.ContextFromHeader(ctx, r)
		defer cancel()

		Err := struct{ Error string }{}
		resEncoder := s.Codec.NewEncoder(w)
		writeError := func(err error) {
			Err.Error = err.Error()
			if e := resEncoder.Encode(Err); e != nil {
				panic(e.Error())
			}
		}

		decoder := s.Codec.NewDecoder(r.Body)
		req, err := svc.DecodeRequest(decoder)
		if err != nil {
			writeError(err)
			return
		}

		res, err := ep(ctx, req)
		if err != nil {
			writeError(err)
			return
		}

		if err := resEncoder.Encode(res); err != nil {
			writeError(err)
		}
	}
}

func (s *Server) AddRoute(svc Service, handlerFunc http.HandlerFunc) {
	s.Mux.HandleFunc("/"+svc.Name(), handlerFunc)
}
