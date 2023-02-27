/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create new task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func add() {
	tasksFile, err := os.Open("./tasks.json")
	if err != nil {
		log.Fatalf("Couldn't open the file: %v", err)
	}
	defer tasksFile.Close()

}
