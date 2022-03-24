package main

import (
	"github.com/kataras/iris/v12"
	"tasks/internal/handlers"
)

func main() {
	app := iris.New()
	server := handlers.NewTaskServer()
	app.Get("/", server.HandlerTask)

	app.Listen(":8080")

}
