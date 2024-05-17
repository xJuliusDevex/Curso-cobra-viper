package models

import (
	"context"
	"fmt"
	"pkg/db"
	"time"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Autor   string             `bson:"autor"`
	Content string             `bson:"content"`
	Done    bool               `bson:"done"`
	Date    time.Time          `bson:"date"`
}

func CreateTask(t Task) (resultOne *mongo.InsertOneResult, err error) {
	tasks, err := db.GetCollection("task")
	if err != nil {
		return nil, err
	}
	result, err := tasks.InsertOne(context.Background(), t)
	if err != nil {
		return nil, err
	}
	defer tasks.Database().Client().Disconnect(context.Background())
	return result, nil
}
func (t Task) String() string {
	done := color.New(color.FgGreen).SprintFunc()
	if !t.Done {
		done = color.New(color.FgRed).SprintFunc()
	}
	green := color.New(color.FgGreen).SprintFunc()

	return fmt.Sprintf("ID: %-20s, \t Authot: %s, \t Task: %-20s , \t Done: %s", t.Id, green((t.Autor)), t.Content, done(t.Done))
}
