package main

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris/v12"
	"log"
	"tasks/internal/handlers"
)

func main() {
	db, err := sqlx.Connect("mysql", "mailtrain:mailtrain@tcp(127.0.0.1:4306)/jackie")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	app := iris.New()
	app.Use(tokenAuth)
	server := handlers.NewTaskServer(db)
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
