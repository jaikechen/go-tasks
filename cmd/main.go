package main

import (
	"errors"
	"github.com/kataras/iris/v12"
	"tasks/internal/handlers"
)

func main() {
	app := iris.New()
	app.Use(tokenAuth)
	server := handlers.NewTaskServer()
	app.Get("/", server.HandlerTask)

	app.Listen(":8080")

}
func tokenAuth(ctx iris.Context) {
	token := ctx.URLParam("token")
	if token != "1234" {
		ctx.StopWithError(403, errors.New("not authorized"))
		return
	}
	ctx.Next()
}
