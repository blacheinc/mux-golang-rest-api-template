package base

import (
	"net/http"

	"github.com/opensaucerer/gotemplate/helper"
	"github.com/opensaucerer/gotemplate/typing"
)

//HealthCheck checks health information of the app in relation to a specific version
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	info := typing.Health{
		Name:    "Golang Server Template",
		Status:  true,
		Version: "0.0.1",
	}

	helper.SendJSONResponse(w, true, http.StatusOK, "Health", typing.M{"health info": info})
}
