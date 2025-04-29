package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gabrielteiga/sarc-pucrs/internal/service"
)

func GetSchedule(w http.ResponseWriter, r *http.Request) {
	param := r.URL.RawQuery

	schedule := service.GetScheduleByURL(param)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&schedule)
	log.Println("final")
}
