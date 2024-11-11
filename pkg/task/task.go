package task

type Task struct {
    ID          int
    Description string
    Priority    string
    Completed   bool
}

func New(description, priority string) Task {
    return Task{
        Description: description,
        Priority:    priority,
        Completed:   false,
    }
}
