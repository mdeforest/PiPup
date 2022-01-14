package router

import (
	"github.com/gorilla/mux"
	"github.com/mdeforest/PiPup/server/internal/app/dispenser/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/start", middleware.StartGame).Methods("POST", "OPTIONS")
	router.HandleFunc("/stop", middleware.StopGame).Methods("POST", "OPTIONS")

	return router
}
