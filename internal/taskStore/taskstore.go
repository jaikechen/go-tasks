package taskStore

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
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
	var tasks []model.Task

	db, err := sql.Open("mysql", "mailtrain:mailtrain@tcp(127.0.0.1:4306)/jackie")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT id,name FROM tasks")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var task model.Task
		err := res.Scan(&task.Id, &task.Text)

		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}
	return tasks
}
