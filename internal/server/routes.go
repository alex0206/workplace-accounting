package server

import (
	"net/http"

	"github.com/alex0206/workplace-accounting/internal/pg"
	"github.com/alex0206/workplace-accounting/internal/server/handlers"
	"github.com/gorilla/mux"
)

// NewAPIRouter getting api handler
func NewAPIRouter(dbConn *pg.DB) http.Handler {
	r := mux.NewRouter()

	f := handlers.NewFactory(dbConn)
	wh := f.WorkPlaceHandler()
	r.HandleFunc("/workplaces", wh.Add).Methods("POST")
	r.HandleFunc("/workplaces/{id:[0-9]+}", wh.Delete).Methods("DELETE")

	return r
}
