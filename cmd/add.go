/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"pkg/models"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Agrega tarea al sistema",
	Long:  `Agrega una tarea al sistema`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Author", viper.GetString("autor"))
		fmt.Println("Task", viper.GetString("task"))
		task := models.Task{
			Autor:   viper.GetString("autor"),
			Content: viper.GetString("task"),
			Done:    false,
			Date:    time.Now(),
		}
		resultOne, err := models.CreateTask(task)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("_ID", resultOne.InsertedID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("task", "t", "tu tarea:", "Tarea que se agrega a la lista")
	addCmd.MarkFlagRequired("task")
	viper.BindPFlag("task", addCmd.Flags().Lookup("task"))

	addCmd.Flags().StringP("autor", "a", "tu nombre:", "Autor que inserto la tarea")
	addCmd.MarkFlagRequired("autor")
	viper.BindPFlag("autor", addCmd.Flags().Lookup("autor"))
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
