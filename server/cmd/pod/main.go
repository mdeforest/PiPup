package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/mdeforest/PiPup/server/internal/app/pod"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

func readConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/pipup/pod")
	viper.AddConfigPath("$HOME/.pipup/pod")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	readConfig()

	app := pod.PodApp{Port: "8080"}
	app.WaitGroup = &sync.WaitGroup{}

	app.WaitGroup.Add(1)

	srv := pod.StartHttpServer(&app)

	if err := srv.Shutdown(context.TODO()); err != nil {
		log.Panic(err)
	}

	app.WaitGroup.Wait()
}
