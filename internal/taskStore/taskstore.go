package taskStore

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	models "tasks/my_models"
	"time"
)

type TaskStore struct {
	db *sqlx.DB
}

func New(d *sqlx.DB) *TaskStore {
	return &TaskStore{
		db: d,
	}
}

func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int {
	return 0
}
func (ts *TaskStore) GetAllTasks(c context.Context) ([]*models.Task, error) {
	return models.Tasks().All(c, ts.db)
}
