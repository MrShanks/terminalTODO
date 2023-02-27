/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	task "github.com/MrShanks/terminalTODO/pkg"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of the tasks",
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list() {
	tasksFile, err := os.Open("./tasks.json")
	if err != nil {
		log.Fatalf("Couldn't open the file: %v", err)
	}
	defer tasksFile.Close()

	byteValue, err := io.ReadAll(tasksFile)
	if err != nil {
		log.Fatalf("Couldn't convert file into a bytes: %v", err)
	}

	var tasks task.Tasks

	json.Unmarshal(byteValue, &tasks)

	for _, task := range tasks.Tasks {
		fmt.Printf("Name: %v\nDescription: %v\n\n", task.Name, task.Desc)
	}
}
