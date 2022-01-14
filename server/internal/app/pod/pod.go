package pod

import (
	"net/http"

	"github.com/mdeforest/PiPup/server/internal/app/pod/router"
	log "github.com/sirupsen/logrus"
)

type PodApp struct {
	Port string

	DispenseEndpoint string
}

func ListenAndServe(a *PodApp) {
	r := router.Router()

	log.Info(("Starting server on the port " + a.Port + "..."))

	log.Fatal(http.ListenAndServe(":"+a.Port, r))
}
