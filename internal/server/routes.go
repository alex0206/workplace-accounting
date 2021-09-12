package server

import (
	"net/http"

	"github.com/alex0206/workplace-accounting/internal/server/handlers/workplace"

	"github.com/gorilla/mux"
)

func Router() http.Handler {
	r := mux.NewRouter()

	wh := workplace.NewHandler()
	r.HandleFunc("/workplaces", wh.Get).Methods("GET")

	return r
}
