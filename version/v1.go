package version

import (
	mux "github.com/gorilla/mux"
	"github.com/opensaucerer/gotemplate/route"
	v1 "github.com/opensaucerer/gotemplate/version/v1"
)

// Version1Routes registers all routes for the v1 version
func Version1Routes(r *mux.Router) {

	// home does not need versioning
	route.RegisterHomeRoutes(r)

	// V1 routes
	router := r.PathPrefix("/v1").Subrouter()

	v1.RegisterBaseRoutes(router)

}
