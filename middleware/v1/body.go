package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/opensaucerer/gotemplate/helper"
	"github.com/opensaucerer/gotemplate/typing"
)

//BodyH is a middleware that parses the request body into the given struct
//and append the decoded data into the request context.
func BodyH(bodyStruct interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//parse body
			err := json.NewDecoder(r.Body).Decode(bodyStruct)
			if err != nil {
				if err.Error() != "EOF" {
					helper.SendJSONResponse(w, false, http.StatusBadRequest, "Error parsing body: "+err.Error(), nil)
					return
				}
			}

			//append decoded data to request context
			ctx := context.WithValue(r.Context(), typing.BodyCtxKey{}, bodyStruct)

			//call next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

//BodyH is a middleware that parses the request body into the given struct
//and append the decoded data into the request context.

func BodyF(bodyStruct interface{}) func(func(http.ResponseWriter, *http.Request)) http.Handler {
	return func(next func(http.ResponseWriter, *http.Request)) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//parse body
			err := json.NewDecoder(r.Body).Decode(bodyStruct)
			if err != nil {
				if err.Error() != "EOF" {
					helper.SendJSONResponse(w, false, http.StatusBadRequest, "Error parsing body: "+err.Error(), nil)
					return
				}
			}
			//append decoded data to request context
			ctx := context.WithValue(r.Context(), typing.BodyCtxKey{}, bodyStruct)

			//call next handler
			next(w, r.WithContext(ctx))
		})
	}

}
