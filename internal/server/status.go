package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
)

func NewStatusEndpoint(c *cache.Cache) http.Handler {
	return &statusEndpoint{c}
}

type statusEndpoint struct {
	cache *cache.Cache
}

func (s *statusEndpoint) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	token := vars["token"]

	if dets, ok := s.cache.Get(token); !ok {
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
}
