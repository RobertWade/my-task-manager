package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "github.com/manifoldco/promptui"
    "github.com/robertwade/my-task-manager/pkg/storage"
    "github.com/robertwade/my-task-manager/pkg/task"
)

func main() {
    for {
        fmt.Println("\n1. Add Task 2. List Tasks 3. Complete Task 4. Delete Task 5. Exit")
        choice, err := promptSelect([]string{"Add Task", "List Tasks", "Complete Task", "Delete Task", "Exit"})
        if err != nil {
            fmt.Printf("Prompt failed %v\n", err)
            continue
        }

        switch choice {
        case "Add Task":
            addTask()
        case "List Tasks":
            listTasks()
        case "Complete Task":
            completeTask()
        case "Delete Task":
            deleteTask()
        case "Exit":
            fmt.Println("Exiting...")
            os.Exit(0)
        }
    }
}

func promptSelect(options []string) (string, error) {
    prompt := promptui.Select{
        Label: "Choose an option",
        Items: options,
    }
    _, result, err := prompt.Run()
    return result, err
}

func addTask() {
    prompt := promptui.Prompt{
        Label: "Enter description",
    }
    desc, _ := prompt.Run()

    priority, err := selectPriority()
    if err != nil {
        fmt.Println("Error selecting priority: ", err)
        return
    }

    newTask := task.New(desc, priority)
    err = storage.AddTask(newTask)
    if err != nil {
        fmt.Printf("Error adding task: %s\n", err)
    } else {
        fmt.Println("Task added!")
    }
}

func selectPriority() (string, error) {
    prompt := promptui.Select{
        Label: "Select Priority",
        Items: []string{"low", "medium", "high"},
    }

    _, result, err := prompt.Run()
    return result, err
}

func listTasks() {
    prompt := promptui.Select{
        Label: "List tasks by priority or all?",
        Items: []string{"All", "Low", "Medium", "High"},
    }

    _, result, err := prompt.Run()
    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        return
    }

    tasks, err := storage.ListTasks()
    if err != nil {
        fmt.Printf("Error retrieving tasks: %s\n", err)
        return
    }

    fmt.Println("\nTasks:")
    for _, t := range tasks {
        if result == "All" || strings.ToLower(t.Priority) == strings.ToLower(result) {
            fmt.Printf("ID: %d, Description: %s, Priority: %s, Completed: %t\n", t.ID, t.Description, t.Priority, t.Completed)
        }
    }
}

func completeTask() {
    fmt.Print("Enter task ID to mark as completed: ")
    input, err := readLine()
    if err != nil {
        fmt.Printf("Error reading input: %s\n", err)
        return
    }
    id, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println("Invalid input. Please enter a valid number.")
        return
    }

    err = storage.CompleteTask(id)
    if err != nil {
        fmt.Printf("Error completing task: %s\n", err)
    } else {
        fmt.Println("Task completed successfully!")
    }
}

func deleteTask() {
    fmt.Print("Enter task ID to delete: ")
    input, err := readLine()
    if err != nil {
        fmt.Printf("Error reading input: %s\n", err)
        return
    }
    id, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println("Invalid input. Please enter a valid number.")
        return
    }

    err = storage.DeleteTask(id)
    if err != nil {
        fmt.Printf("Error deleting task: %s\n", err)
    } else {
        fmt.Println("Task deleted successfully!")
    }
}

func readLine() (string, error) {
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(input), nil
}
