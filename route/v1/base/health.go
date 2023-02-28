package base

import (
	"github.com/opensaucerer/gotemplate/controller/v1/base"

	mux "github.com/gorilla/mux"
)

func RegisterHealthRoutes(r *mux.Router) {

	router := r.PathPrefix("/health").Subrouter()

	router.HandleFunc("", base.HealthCheck).Methods("GET")
}
