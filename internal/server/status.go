package server

import (
	"encoding/json"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"net/http"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

const (
	tokenName = "token"
)

const StatusEndpointPath = "/process/dataset/{" + tokenName + "}/status"

func NewStatusEndpoint(c *cache.Cache) middle.LoggedMiddlewareFn {
	return func(log *logrus.Entry) midl.Middleware {
		return &statusEndpoint{c, log}
	}
}

type statusEndpoint struct {
	cache *cache.Cache
	log *logrus.Entry
}

func (s *statusEndpoint) Handle(req midl.Request) midl.Response {
	vars := mux.Vars(req.RawRequest())
	token := vars[tokenName]

	if dets, ok := s.cache.Get(token); !ok {
		logrus.
			res.WriteHeader(http.StatusNotFound)

		// TODO: log this
	} else {
		if bytes, err := json.Marshal(dets); err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			// TODO: Log this
		} else {
			_, err = res.Write(bytes)
			// TODO: log error
		}
	}
	panic("implement me")
}

func (s *statusEndpoint) ServeHTTP(res http.ResponseWriter, req *http.Request) {

}
