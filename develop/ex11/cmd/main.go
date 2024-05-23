package main

import (
	"ex11/internal/handler"
	"ex11/internal/server"
	"ex11/internal/service"
	"ex11/internal/storage"
)

func main() {
	st := storage.NewStorage()
	s := service.NewService(st)
	h := handler.NewHandler(s)

	app := server.NewServer(h.InitRoutes())

	app.Run()
}
