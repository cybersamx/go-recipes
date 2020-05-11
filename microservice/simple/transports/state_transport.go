package transports

import (
	"context"
	"github.com/cybersamx/go-recipes/microservice/simple/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
	"path"
)

func DecodeGetStatesRequest() httptransport.DecodeRequestFunc {
	return func (_ context.Context, r *http.Request) (interface{}, error) {
		return nil, nil
	}
}

func DecodeGetStateRequest() httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var req endpoints.GetStateRequest
		req.Abbreviation = path.Base(r.URL.Path)
		return req, nil
	}
}
