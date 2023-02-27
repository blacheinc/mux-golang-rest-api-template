package base

import (
	"net/http"

	"github.com/opensaucerer/gotemplate/helper"
	"github.com/opensaucerer/gotemplate/typing"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	data := typing.Health{
		Status:      true,
		Version:     "0.0.1",
		Description: "Golang Server Template",
	}
	helper.SendJSONResponse(w, true, http.StatusOK, "Health Check", typing.M{"health": data})
}

func HealthInformation(w http.ResponseWriter, r *http.Request) {
	info := typing.HealthInfo{
		Status: "OK",
	}
	helper.SendJSONResponse(w, true, http.StatusOK, "Health Information", typing.M{"health info": info})
}
