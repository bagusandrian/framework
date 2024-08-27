package main

import (
	"log"

	"github.com/bagusandrian/framework/internal/config"
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
	log.Default().Println("Starting dummy http")

	// init config
	cfg, err := config.New(repoName, config.WithFeatureFlag())
	if err != nil {
		log.Fatalf("failed to init the config: %v", err)
	}
	log.Println("init config done")

	// set panic option
	// if !env.IsDevelopment() {
	// 	svcEnv := string(env.ServiceEnv())
	// 	panics.SetOptions(&panics.Options{
	// 		Env:             svcEnv,
	// 		SlackWebhookURL: cfg.SlackWebhook,
	// 		SlackChannel:    "product-review-alert",
	// 		CustomMessage:   "@morty-dev",
	// 		Tags:            map[string]string{"service": appName, "env": svcEnv, "hostname": cfg.Server.HostName, "hostip": cfg.Server.HostIP},
	// 	})
	// }

	// init metrics & callwrapper
	// command to move line

	err = startApp(cfg)
	if err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
