package db_test

import (
	"context"
	"fmt"
	"pkg/db"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestGetConn(t *testing.T) {
	client, err := db.GetConnection()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	listdb, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	defer client.Disconnect(context.Background())
	fmt.Println(listdb)
}
func TestGetCollection(t *testing.T) {
	tasks, err := db.GetCollection("task")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	result := tasks.FindOne(context.Background(), bson.D{})
	var doc bson.D
	err = result.Decode(&doc)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	tasks.Database().Client().Disconnect(context.Background())
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	fmt.Println(doc)
}
func TestConfig(t *testing.T) {
	db.LoadConfig()
}
