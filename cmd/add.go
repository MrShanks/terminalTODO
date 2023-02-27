package cmd

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	task "github.com/MrShanks/terminalTODO/pkg"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create new task",
	Run: func(cmd *cobra.Command, args []string) {
		add(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func add(input []string) {
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

	parsedInput := strings.Join(input, " ")

	rawTask := strings.Split(parsedInput, ":")

	task := task.NewTask(rawTask[0], rawTask[1])

	tasks.Tasks = append(tasks.Tasks, task)

	jsonTasks, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		log.Fatalf("Couldn't marshal tasks: %v", err)
	}

	os.WriteFile("./tasks.json", jsonTasks, 0644)
}
