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
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Elimina tareas de la base de datos",
	Long:  `Elimina tareas de la base de datos`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
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
		result, err := tasks.DeleteMany(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
