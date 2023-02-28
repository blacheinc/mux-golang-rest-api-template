package base

import (
	"github.com/opensaucerer/gotemplate/controller/v1/base"

	mux "github.com/gorilla/mux"
)

func RegisterHomeRoute(r *mux.Router) {

	r.HandleFunc("/", base.Home).Methods("GET")
}
