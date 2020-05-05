package upload

import (
	"fmt"
	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
)

const (
	errBadToken = "Invalid UUID format for process id"
	errNotFound = "No prepared dataset found for process id '%s'.  Either the " +
		"timed out or was not posted."
)

// NewTokenFilter returns a middleware layer that validates
// that the token is a valid format and exists in the
// metadata cache.
func NewTokenFilter(c *cache.Cache) middle.LoggedMiddlewareFn {
	return func(logger *logrus.Entry) midl.Middleware {
		return midl.MiddlewareFunc(func(req midl.Request) midl.Response {
			val := mux.Vars(req.RawRequest())
			tkn := val[tokenKey]

			if _, err := uuid.Parse(tkn); err != nil {
				return svc.BadRequest(errBadToken)
			}

			if _, ok := c.Get(tkn); !ok {
				return svc.NotFound(fmt.Sprintf(errNotFound, tkn))
			}

			return nil
		})
	}
}
