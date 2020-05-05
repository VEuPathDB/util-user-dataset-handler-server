package metadata

import (
	// Std Lib
	"encoding/json"
	"net/http"

	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
)

const (
	errEmptyMetadata = "metadata payload cannot be empty"
	errParseMetadata = "failed to parse input JSON: %s"
)

// NewMetadataValidator is a validation wrapper middleware
// that attempts to parse and validate the request body as
// a metadata JSON payload, either calling the endpoint on
// success or returning an error response to the caller.
func NewMetadataValidator(next func(meta *Metadata) midl.Middleware) midl.MiddlewareFunc {
	return func(req midl.Request) midl.Response {
		if data, err := parseMetadata(req); err != nil {
			return err
		} else {
			return next(data).Handle(req)
		}
	}
}

func parseMetadata(req midl.Request) (*Metadata, midl.Response) {
	bytes := req.Body()
	if req.Body() == nil {
		return nil, midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
			Status:  svc.StatusBadRequest,
			Message: errEmptyMetadata,
		})
	}

	var foo Metadata
	if err := json.Unmarshal(bytes, &foo); err != nil {
		return nil, midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
			Status:  svc.StatusBadRequest,
			Message: errParseMetadata,
		})
	}

	if err := validateMetadata(&foo); err != nil {
		return nil, err
	}

	return &foo, nil
}

func validateMetadata(meta *Metadata) midl.Response {
	if val := meta.Validate(); !val.Ok {
		return midl.MakeResponse(http.StatusUnprocessableEntity,
			&svc.ValidationResponse{
				Status: svc.StatusBadInput,
				Reasons: val.Result,
			})
	}
	return nil
}