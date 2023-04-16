package controller

import (
	"net/http"

	"github.com/blacheinc/gotemplate/helper"
	"github.com/blacheinc/gotemplate/typing"
)

func Home(w http.ResponseWriter, r *http.Request) {

	data := typing.Home{
		Status:      true,
		Version:     "1.0.0",
		Description: "Golang Server Template",
	}

	helper.SendJSONResponse(w, true, http.StatusOK, "Home", typing.M{"home": data})
}
