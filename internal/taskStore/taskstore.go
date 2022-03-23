package taskStore

import (
	"tasks/model"
	"time"
)

type TaskStore struct {
	tasks  map[int]model.Task
	nextId int
}

func New() *TaskStore {
	return &TaskStore{}
}

func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int {
	ts.nextId++
	t := model.Task{
		Id:   ts.nextId,
		Text: text,
		Tags: tags,
		Due:  due,
	}
	t.Due.Date()
	ts.tasks[ts.nextId] = t
	return ts.nextId
}
func (ts *TaskStore) GetAllTasks() []model.Task {
	return []model.Task{
		{
			Id:   1,
			Text: "task1",
		},
		{
			Id:   2,
			Text: "task2",
		},
	}
}
