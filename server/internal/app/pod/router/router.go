package router

import (
	"sync"

	"github.com/gorilla/mux"
	"github.com/mdeforest/PiPup/server/internal/app/pod/middleware"
)

func Router(wg *sync.WaitGroup) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/start", middleware.StartGame(wg)).Methods("POST", "OPTIONS")
	router.HandleFunc("/stop", middleware.StopGame(wg)).Methods("POST", "OPTIONS")

	return router
}
