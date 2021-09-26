package main

import (
	"github.com/gribanoid/go-rest-api"
	"github.com/gribanoid/go-rest-api/pkg/handeler"
	"log"
)

func main()  {
	handlers := new(handeler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occuarid while running http server: %s", err.Error())
	}
}