package task

import "time"

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Status   string `json:"status"`
	Assignee string `json:"assignee"`
	metadata `json:"metadata"`
}

type metadata struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTask(name, desc string) Task {
	return Task{
		Name:   name,
		Desc:   desc,
		Status: "TODO",
		metadata: metadata{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func (t *Task) GetStatus() string {
	return t.Status
}
