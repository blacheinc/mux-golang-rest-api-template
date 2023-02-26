package base

import (
	"net/http"

	"github.com/opensaucerer/gotemplate/helper"
	"github.com/opensaucerer/gotemplate/typing"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	data := typing.Health{
		Version:     "0.0.1",
		Status:      true,
		Description: "Raid2Earn. Leaders create raids and pawns pick 'em up.",
	}
	helper.SendJSONResponse(w, true, http.StatusOK, "Health Check", typing.M{"health": data})
}
