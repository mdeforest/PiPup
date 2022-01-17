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

func ListenAndServe(a *PodApp) {
	a.WaitGroup = &sync.WaitGroup{}

	r := router.Router(a.WaitGroup)

	log.Info(("Starting server on the port " + a.Port + "..."))

	a.WaitGroup.Add(1)

	go func() {
		defer a.WaitGroup.Done()

		if err := http.ListenAndServe(":"+a.Port, r); err != http.ErrServerClosed {
			log.Fatal("ListenAndServe(): %v", err)
		}
	}()

	a.WaitGroup.Wait()
}
