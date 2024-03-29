package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/blacheinc/gotemplate/helper"
	"github.com/blacheinc/gotemplate/typing"
)

// Body parses the request body into the given struct
func Body(bodyStruct interface{}, w http.ResponseWriter, r *http.Request) context.Context {
	//parse body
	err := json.NewDecoder(r.Body).Decode(bodyStruct)
	if err != nil {
		if err.Error() != "EOF" {
			helper.SendJSONResponse(w, false, http.StatusBadRequest, "Error parsing body: "+err.Error(), nil)
			return nil
		}
	}
	//append decoded data to request context
	return context.WithValue(r.Context(), typing.BodyCtxKey{}, bodyStruct)
}

// BodyH is a middleware that takes a struct and a request handler.
// It parses the request body into the given struct and pushes it into
// the request's context using the BodyCtxKey
func BodyH(bodyStruct interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//parse body
			ctx := Body(bodyStruct, w, r)
			if ctx == nil {
				return
			}
			//call next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// BodyF is a middleware that takes a struct and a controller function.
// It parses the request body into the given struct and pushes it into
// the request's context using the BodyCtxKey
func BodyF(bodyStruct interface{}) func(func(http.ResponseWriter, *http.Request)) http.Handler {
	return func(next func(http.ResponseWriter, *http.Request)) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//parse body
			ctx := Body(bodyStruct, w, r)
			if ctx == nil {
				return
			}
			//call next handler
			next(w, r.WithContext(ctx))
		})
	}
}
