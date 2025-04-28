package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielteiga/sarc-pucrs/internal/service"
)

type Response struct {
	ID       string `json:"id"`
	Schedule string `json:"schedule"`
}

func GetSchedule(w http.ResponseWriter, r *http.Request) {
	param := r.URL.RawQuery
	println(param)
	service.GetScheduleByURL(param)

	example := &Response{}
	json.NewEncoder(w).Encode(&example)
}
