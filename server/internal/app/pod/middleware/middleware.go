package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/mdeforest/PiPup/server/pkg/games/pod"
	httpresponses "github.com/mdeforest/PiPup/server/pkg/httpResponses"
)

func StartGame(wg *sync.WaitGroup) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		level := params.Get("level")
		length := params.Get("length")

		if level == "" {
			w.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(w).Encode(httpresponses.ErrorResponse("level not included"))

			return
		}

		if length == "" {
			w.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(w).Encode(httpresponses.ErrorResponse("time length not included"))

			return
		}

		levelInt, err := strconv.Atoi(level)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(httpresponses.ErrorResponse("level is not a valid value"))
		}

		lengthInt, err := strconv.Atoi(length)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(httpresponses.ErrorResponse("length is not a valid value"))
		}

		game := pod.NewPodGame(levelInt, lengthInt)

		wg.Add(1)

		go game.Start(wg)

		w.Header().Set("Content-Type", "application.json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(httpresponses.OkResponse())
	}
}

func StopGame(wg *sync.WaitGroup) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
