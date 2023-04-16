package v1

import (
	"github.com/blacheinc/gotemplate/route/v1/base"
	"github.com/gorilla/mux"
)

func RegisterBaseRoutes(r *mux.Router) {

	router := r.PathPrefix("/").Subrouter()

	base.RegisterHealthRoutes(router)
}
