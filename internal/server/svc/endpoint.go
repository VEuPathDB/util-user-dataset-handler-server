package svc

import "github.com/gorilla/mux"

type Endpoint interface {
	Register(r *mux.Router)
}
