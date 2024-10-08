package main

import (
	"log"

	"github.com/bagusandrian/framework/internals/config"
)

var cfgTest bool

const (
	repoName = "framework"
	appName  = repoName + "-http"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("%v", err)
		}
	}()
	log.Println("Starting dummy http")

	// init config
	cfg, err := config.New(repoName)
	if err != nil {
		log.Fatalf("failed to init the config: %v", err)
	}
	log.Println("init config done")

	err = startApp(cfg)
	if err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
