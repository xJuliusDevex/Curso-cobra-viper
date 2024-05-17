/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"pkg/db"
	"pkg/models"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista las tareas registrada",
	Long:  `Lista las tareas registrada`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list task")
		task, err := db.GetCollection("task")
		if err != nil {
			log.Fatal(err)
		}
		curs, err := task.Find(context.Background(), bson.D{{"done", viper.GetBool("done")}})
		if err != nil {
			log.Fatal(err)
		}
		defer curs.Close(context.Background())
		for curs.Next(context.Background()) {
			var task models.Task
			err = curs.Decode(&task)
			if err != nil {
				log.Fatal(err)
			}
			color.New().Println(task)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("done", "d", false, "muestra las tareas no compleatadas")
	viper.BindPFlag("done", listCmd.Flags().Lookup("done"))
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
