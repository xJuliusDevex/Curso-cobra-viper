package models_test

import (
	"fmt"
	"pkg/models"
	"testing"
	"time"
)

func TestCreateTask(t *testing.T) {
	task1 := models.Task{
		Autor:   "Julius",
		Content: "Limpiar",
		Done:    false,
		Date:    time.Now(),
	}
	re, err := models.CreateTask(task1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	fmt.Println("_id", re.InsertedID)
}
