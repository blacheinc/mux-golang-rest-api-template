package middleware

import (
	"context"
	"net/http"

	"github.com/blacheinc/gotemplate/typing"
	"github.com/gorilla/mux"
)

// ParamsH is a middleware that takes a request handler
// and pushes the request path parameters it into the request's context
// using the ParamsCtxKey.
func ParamsH(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//set path parameter to the request context
		//call next handler
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), typing.ParamsCtxKey{}, mux.Vars(r))))

	})
}

// ParamsF is a middleware that takes a controller function
// and pushes the request path parameters into the request's context
// using the ParamsCtxKey.
func ParamsF(next func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//set path parameter to the request context
		//call next handler
		next(w, r.WithContext(context.WithValue(r.Context(), typing.ParamsCtxKey{}, mux.Vars(r))))
	})
}
