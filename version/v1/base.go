package v1

import (
	"github.com/gorilla/mux"
	"github.com/opensaucerer/gotemplate/route/v1/base"
)

func RegisterBaseRoutes(r *mux.Router) {

	router := r.PathPrefix("/").Subrouter()

	base.RegisterHealthRoutes(router)
}
