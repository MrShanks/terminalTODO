package task

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
