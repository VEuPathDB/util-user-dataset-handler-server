package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"

	"github.com/VEuPathDB/util-exporter-server/internal/server/xio"
)

const (
	errEmptyMetadata = "metadata payload cannot be empty"
	errParseMetadata = "failed to parse input JSON: %s"
)

const (
	_400 = http.StatusBadRequest
	_422 = http.StatusUnprocessableEntity
)

func NewMetadataWrapper(next func(meta *xio.Metadata) midl.Middleware) midl.MiddlewareFunc {
	return func(req midl.Request) midl.Response {
		if data, err := parseMetadata(req); err != nil {
			return err
		} else {
			data
		}
	}
}

func parseMetadata(req midl.Request) (*xio.Metadata, midl.Response) {
	bytes := req.Body()
	if req.Body() == nil {
		return nil, midl.MakeErrorResponse(_400, errors.New(errEmptyMetadata))
	}

	var foo xio.Metadata
	if err := json.Unmarshal(bytes, &foo); err != nil {
		return nil, midl.MakeErrorResponse(_400, fmt.Errorf(errParseMetadata, err))
	}

	if err := validateMetadata(&foo); err != nil {
		return nil, midl.MakeResponse()
	}
}

func validateMetadata(meta *xio.Metadata) midl.Response {
	errs := make(map[string][]string)
	if meta.Owner == 0 {
		errs["owner"] = []string{"owner must be set"}
	}

	if len(meta.Projects) == 0 {
		errs["projects"] = []string{"at least one project must be provided"}
	} else {
		for _, pro := range meta.Projects {
			if pro.IsValid() {
				continue
			}
			if _, ok := errs["projects"]; !ok {
				errs["projects"] = []string{}
			}
			errs["projects"] = append(errs["projects"],
				fmt.Sprintf("unrecognized project '%s'", pro))
		}
	}

	if len(meta.Type.Version) == 0 {
		errs["type.version"] = []string{"type.version is required"}
	}
	if len(meta.Type.Name) == 0 {
		errs["type.name"] = []string{"type.name is required"}
	}
	if len(meta.Token) == 0 {
		errs["token"] = []string{"token is required"}
	} else {
		_, err := uuid.Parse(meta.Token)
		if err != nil {
			errs["token"] = []string{"token must be a valid uuid"}
		}
	}
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
