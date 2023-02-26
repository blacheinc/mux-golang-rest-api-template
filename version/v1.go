package version

import (
	mux "github.com/gorilla/mux"
	"github.com/opensaucerer/gotemplate/route/v1/base"
)

// Version1Routes registers all routes for the v1 version
func Version1Routes(r *mux.Router) {

	// this doesn't need versioning, yet
	base.RegisterHealthRoutes(r)

	// // V1 routes
	// router := r.PathPrefix("/v1").Subrouter()

	// base.RegisterPricingRoutes(router)
	// user.RegisterAuthRoutes(router)
	// user.RegisterUserRoutes(router)
	// user.RegisterActionRoutes(router)
	// raid.RegisterSimulateRoutes(router)
	// raid.RegisterRaidRoutes(router)

}
