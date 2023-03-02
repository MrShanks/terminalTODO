package task

import (
	"fmt"
	"testing"
)

func TestTaskCreation(t *testing.T) {
	var tt = []struct {
		name     string
		desc     string
		testDesc string
		expected string
	}{
		{
			name:     "test",
			desc:     "test",
			expected: "Name: test\nDescription: test\nStatus: TODO\n\n",
			testDesc: "Create a new task with name and desc test should return test for the task name and desc",
		},
		{
			name:     "",
			desc:     "",
			expected: "Name: \nDescription: \nStatus: TODO\n\n",
			testDesc: "Create a new task with name and desc empty should create a task with name and desc empty",
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf(tc.testDesc), func(t *testing.T) {
			task := NewTask(tc.name, tc.desc)
			taskString := fmt.Sprintf("Name: %v\nDescription: %v\nStatus: %v\n\n", task.Name, task.Desc, task.GetStatus())
			if taskString != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, taskString)
			}
		})
	}
}
