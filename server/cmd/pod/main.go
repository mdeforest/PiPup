package main

import (
	"fmt"

	"github.com/mdeforest/PiPup/server/internal/app/pod"
	"github.com/spf13/viper"
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

	pod.ListenAndServe(&app)
}
