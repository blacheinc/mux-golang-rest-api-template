package version

import (
	"github.com/blacheinc/gotemplate/route"
	v1 "github.com/blacheinc/gotemplate/version/v1"
	mux "github.com/gorilla/mux"
)

// Version1Routes registers all routes for the v1 version
func Version1Routes(r *mux.Router) {

	// home does not need versioning
	route.RegisterHomeRoutes(r)

	// V1 routes
	router := r.PathPrefix("/v1").Subrouter()

	v1.RegisterBaseRoutes(router)

}
