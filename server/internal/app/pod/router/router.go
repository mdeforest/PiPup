package router

import (
	"github.com/gorilla/mux"
	"github.com/mdeforest/PiPup/server/internal/app/pod/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/start", middleware.StartGame).Methods("POST", "OPTIONS")
	router.HandleFunc("/stop", middleware.StopGame).Methods("POST", "OPTIONS")
	router.HandleFunc("/test", middleware.Test).Methods("POST", "OPTIONS")

	return router
}
