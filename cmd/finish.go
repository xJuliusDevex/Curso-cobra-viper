/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"pkg/db"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// finishCmd represents the finish command
var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finaliza una tarea",
	Long:  `Finaliza una tarea`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("finish task")
		tasks, err := db.GetCollection("task")
		if err != nil {
			log.Fatal(err)
		}
		defer tasks.Database().Client().Disconnect(context.Background())
		ids, err := IdsToObjectId(args)
		if err != nil {
			log.Fatal(err)
		}
		filter := bson.M{"_id": bson.M{"$in": ids}}
		update := bson.M{"$set": bson.M{"done": true}}
		result, err := tasks.UpdateMany(context.Background(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(finishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// finishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// finishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func IdsToObjectId(ids []string) (olds []primitive.ObjectID, err error) {
	for _, v := range ids {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return nil, err
		}
		olds = append(olds, id)
	}
	return olds, nil
}
