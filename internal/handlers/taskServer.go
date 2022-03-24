package handlers

import (
	"github.com/kataras/iris/v12"
	"tasks/internal/taskStore"
)

type taskServer struct {
	store *taskStore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskStore.New()
	server := taskServer{
		store: store,
	}
	return &server
}

func (t taskServer) HandlerTask(ctx iris.Context) {
	tasks := t.store.GetAllTasks()
	ctx.JSON(tasks)
}
