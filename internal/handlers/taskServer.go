package handlers

import (
	"encoding/json"
	"net/http"
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

func (t taskServer) HandlerTask(w http.ResponseWriter, req *http.Request) {
	tasks := t.store.GetAllTasks()
	js, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
