package taskStore

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	store := New()
	if store == nil {
		t.Error("create store fail")
	}
}
func TestCreate(t *testing.T) {
	store := New()
	store.CreateTask("studying go", []string{"go"}, time.Now())
}
