package types

import "github.com/gorilla/mux"

// Endpoint defines a standalone HTTP endpoint with a route registration method.
type Endpoint interface {

	// Register is used to register the current Endpoint instance with the given
	// router.
	Register(r *mux.Router)
}
