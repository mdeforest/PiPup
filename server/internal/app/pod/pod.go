package pod

import (
	"net/http"
	"sync"

	"github.com/mdeforest/PiPup/server/internal/app/pod/router"
	log "github.com/sirupsen/logrus"
)

type PodApp struct {
	Port string

	DispenseEndpoint string
	WaitGroup        *sync.WaitGroup
}

func StartHttpServer(a *PodApp) *http.Server {
	srv := &http.Server{Addr: ":" + a.Port}
	r := router.Router(a.WaitGroup)

	log.Info(("Starting server on the port " + a.Port + "..."))

	go func() {
		defer a.WaitGroup.Done()

		if err := http.ListenAndServe(":"+a.Port, r); err != http.ErrServerClosed {
			log.Fatal("ListenAndServe(): %v", err)
		}
	}()

	return srv
}
