package base

import (
	"net/http"

	"github.com/opensaucerer/gotemplate/helper"
	"github.com/opensaucerer/gotemplate/typing"
)

func Home(w http.ResponseWriter, r *http.Request) {

	data := typing.Home{
		Status:      true,
		Version:     "0.0.1",
		Description: "Golang Server Template",
	}

	helper.SendJSONResponse(w, true, http.StatusOK, "Home", typing.M{"health": data})
}
