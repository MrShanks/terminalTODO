package task

import "time"

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Status   status `json:"status"`
	metadata `json:"metadata"`
}

type metadata struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type status int

const (
	TODO status = 1
	WIP
	DONE
)

func NewTask(name, desc string) Task {
	return Task{
		Name:   name,
		Desc:   desc,
		Status: TODO,
		metadata: metadata{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func (t *Task) Progress() {
	if t.Status >= 3 {
		return
	}
	t.Status = t.Status + 1
}

func (t *Task) UnProgress() {
	if t.Status <= 1 {
		return
	}
	t.Status = t.Status - 1
}
