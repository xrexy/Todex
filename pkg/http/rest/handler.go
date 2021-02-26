package rest

import (
	"github.com/gorilla/mux"
)

// InitHandlers is intializing all handlers
func InitHandlers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler())

	router.HandleFunc("/help", welcomeHandler())
	return router
}
