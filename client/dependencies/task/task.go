package task

type Task struct {
	Owner       string
	Name        string
	Description string
	Done        bool
}

func NewTask(owner, name, description string) *Task {
	return &Task{
		Owner:       owner,
		Name:        name,
		Description: description,
		Done:        false,
	}
}

func (t *Task) DoneTask() {
	t.Done = !t.Done
	// t.Done = true
}
