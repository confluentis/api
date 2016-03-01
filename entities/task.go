package entities

type Task struct {
    Id int64    `db:"id" json:"id"`
}

func newTask() Task {
    return Task{}
}