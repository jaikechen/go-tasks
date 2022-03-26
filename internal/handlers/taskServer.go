package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris/v12"
	"tasks/internal/taskStore"
)

type taskServer struct {
	store *taskStore.TaskStore
}

func NewTaskServer(d *sqlx.DB) *taskServer {

	store := taskStore.New(d)
	server := taskServer{
		store: store,
	}
	return &server
}

func (t taskServer) HandlerTask(ctx iris.Context) {
	tasks, err := t.store.GetAllTasks(ctx.Request().Context())
	if err != nil {
		ctx.StopWithError(500, err)
	}
	ctx.JSON(tasks)
}
