package base

import (
	"github.com/blacheinc/gotemplate/controller/v1/base"

	mux "github.com/gorilla/mux"
)

func RegisterHealthRoutes(r *mux.Router) {

	router := r.PathPrefix("/health").Subrouter()

	router.HandleFunc("", base.Health).Methods("GET")
}
