package job

import (
	// Std Lib
	"encoding/json"
	"net/http"

	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
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
func NewMetadataValidator() midl.MiddlewareFunc {
	return func(req midl.Request) midl.Response {
		if data, err := parseMetadata(req); err != nil {
			return err
		} else {
			req.AdditionalContext()["data"] = data
		}

		return nil
	}
}

func parseMetadata(req midl.Request) (*Metadata, midl.Response) {
	log := req.AdditionalContext()[middle.KeyLogger].(*logrus.Entry)

	bytes := req.Body()
	if req.Body() == nil {
		log.WithField("status", http.StatusBadRequest).Info(errEmptyMetadata)
		return nil, midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
			Status:  svc.StatusBadRequest,
			Message: errEmptyMetadata,
		})
	}

	var foo Metadata
	if err := json.Unmarshal(bytes, &foo); err != nil {
		log.WithField("status", http.StatusBadRequest).Info(errParseMetadata)
		return nil, midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
			Status:  svc.StatusBadRequest,
			Message: errParseMetadata,
		})
	}

	if err := validateMetadata(&foo, log); err != nil {
		return nil, err
	}

	return &foo, nil
}

func validateMetadata(meta *Metadata, log *logrus.Entry) midl.Response {
	if val := meta.Validate(); !val.Ok {
		log.WithField("status", http.StatusUnprocessableEntity).
			Info("metadata validation failed")
		return midl.MakeResponse(http.StatusUnprocessableEntity,
			&svc.ValidationResponse{
				Status: svc.StatusBadInput,
				Reasons: val.Result,
			})
	}
	return nil
}