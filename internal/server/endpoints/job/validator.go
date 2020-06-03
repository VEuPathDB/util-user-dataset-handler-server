package job

import (
	// Std Lib
	"encoding/json"
	"fmt"
	"net/http"

	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/service/logger"
)

const (
	errEmptyMetadata = "metadata payload cannot be empty"
	errParseMetadata = "failed to parse input JSON: %s"
	dataCtxKey       = "data"
)

// NewMetadataValidator is a validation wrapper middleware
// that attempts to parse and validate the request body as
// a metadata JSON payload, either calling the endpoint on
// success or returning an error response to the caller.
func NewMetadataValidator() midl.MiddlewareFunc {
	return func(req midl.Request) midl.Response {
		data, err := parseMetadata(req)

		if err != nil {
			return err
		}

		req.AdditionalContext()[dataCtxKey] = data

		return nil
	}
}

func parseMetadata(req midl.Request) (*Metadata, midl.Response) {
	log := logger.ByRequest(req)

	if req.Body() == nil {
		log.WithField("status", http.StatusBadRequest).Info(errEmptyMetadata)

		return nil, midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
			Status:  svc.StatusBadRequest,
			Message: errEmptyMetadata,
		})
	}

	bytes := req.Body()
	log.Debug("Request body: ", string(bytes))

	var foo Metadata
	if err := json.Unmarshal(bytes, &foo); err != nil {
		errTxt := fmt.Sprintf(errParseMetadata, err)

		log.WithField("status", http.StatusBadRequest).Info(errTxt)

		return nil, midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
			Status:  svc.StatusBadRequest,
			Message: errTxt,
		})
	}

	foo.Token = mux.Vars(req.RawRequest())[tokenKey]

	if err := validateMetadata(&foo, log); err != nil {
		return nil, err
	}

	return &foo, nil
}

func validateMetadata(meta *Metadata, log *logrus.Entry) midl.Response {
	if val := meta.Validate(); val.Failed {
		log.WithField("status", http.StatusUnprocessableEntity).
			Info("metadata validation failed")

		return midl.MakeResponse(http.StatusUnprocessableEntity,
			&svc.ValidationResponse{
				Status: svc.StatusBadInput,
				Errors: svc.ValidationBundle{
					ByKey: val.Result,
				},
			})
	}

	return nil
}
