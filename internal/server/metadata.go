package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/VEuPathDB/util-exporter-server/internal/server/xio"
	"github.com/patrickmn/go-cache"
	"gopkg.in/foxcapades/go-midl.v1/pkg/midl"
	"net/http"
)

const (
	errEmptyMetadata  = "metadata payload cannot be empty"
	errParseMetadata  = "failed to parse input JSON: %s"

)

const (
	_400 = http.StatusBadRequest
	_422 = http.StatusUnprocessableEntity
)

func NewMetadataWrapper(next func(meta *xio.Metadata) midl.Middleware) midl.MiddlewareFunc {
	return func(req midl.Request) midl.Response {

		if data, err := parseMetadata(req); err != nil {
			return err
		}
	}
}

func parseMetadata(req midl.Request) (*xio.Metadata, midl.Response) {
	bytes := req.Body()
	if req.Body() == nil {
		return nil, midl.MakeErrorResponse(_400, errors.New(errEmptyMetadata))
	}

	var foo metadata
	if err := json.Unmarshal(bytes, &foo); err != nil {
		return nil, midl.MakeErrorResponse(_400, fmt.Errorf(errParseMetadata, err))
	}

	if err := validateMetadata(&foo)
}

func validateMetadata(meta *xio.Metadata) error {

}

func NewMetadataEndpoint(c *cache.Cache) func(*xio.Metadata) midl.Middleware {
	return func(meta *xio.Metadata) midl.Middleware {
		return &metadataEndpoint{meta}
	}
}

type metadataEndpoint struct {
	meta *xio.Metadata
}

func (m *metadataEndpoint) Handle(req midl.Request) midl.Response {
	panic("implement me")
}
