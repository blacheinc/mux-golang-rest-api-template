package base

import (
	"net/http"

	"github.com/opensaucerer/gotemplate/helper"
	"github.com/opensaucerer/gotemplate/typing"
)

//HealthCheck checks health information of the app in relation to a specific version
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	info := typing.Health{
		Status: "OK",
	}
	helper.SendJSONResponse(w, true, http.StatusOK, "Health Information", typing.M{"health info": info})
}
