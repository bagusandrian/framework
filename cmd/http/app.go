package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bagusandrian/framework/internals/config"
	hImpl "github.com/bagusandrian/framework/internals/handler/http/impl"
	dbImpl "github.com/bagusandrian/framework/internals/repository/db/impl"
	ucImpl "github.com/bagusandrian/framework/internals/usecase/dummy/impl"
	"github.com/gofiber/fiber/v2"
)

func startApp(cfg *config.Config) (err error) {

	log.Println("starting http server", time.Now())
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	repositoryDB := dbImpl.New(cfg)
	usecaseDummy := ucImpl.New(cfg, repositoryDB)
	handlerDummy := hImpl.New(cfg, usecaseDummy)

	app := fiber.New(fiber.Config{
		AppName: "framework",
	})
	app.Post("/", handlerDummy.Dummy)
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", cfg.Server.HTTP.Port)))
	return nil
}
