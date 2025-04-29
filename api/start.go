package api

import (
	"net/http"

	"github.com/gabrielteiga/sarc-pucrs/api/controller"
)

func Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/schedule", controller.GetSchedule)

	http.ListenAndServe(":8080", mux)
}
