package storage

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/robertwade/my-task-manager/pkg/task"
)

const fileName = "tasks.csv"

func init() {
    if _, err := os.Stat(fileName); os.IsNotExist(err) {
        file, err := os.Create(fileName)
        if err != nil {
            panic(fmt.Sprintf("Cannot create file: %s", err))
        }
        defer file.Close()
    }
}

func AddTask(t task.Task) error {
    randomID, err := generateRandomID()
    if err != nil {
        return fmt.Errorf("failed to generate random ID: %v", err)
    }
    t.ID = randomID

    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("failed to open file for writing: %v", err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    record := []string{strconv.Itoa(t.ID), t.Description, t.Priority, strconv.FormatBool(t.Completed)}
    if err := writer.Write(record); err != nil {
        return fmt.Errorf("failed to write record: %v", err)
    }
    return nil
}

func ListTasks() ([]task.Task, error) {
    file, err := os.Open(fileName)
    if err != nil {
        return nil, fmt.Errorf("failed to open file for reading: %v", err)
    }
    defer file.Close()

    var tasks []task.Task
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, fmt.Errorf("failed to read records: %v", err)
    }

    for _, record := range records {
        id, err := strconv.Atoi(record[0])
        if err != nil {
            return nil, fmt.Errorf("failed to convert ID: %v", err)
        }
        completed, err := strconv.ParseBool(record[3])
        if err != nil {
            return nil, fmt.Errorf("failed to parse completed status: %v", err)
        }
        tasks = append(tasks, task.Task{
            ID:          id,
            Description: record[1],
            Priority:    record[2],
            Completed:   completed,
        })
    }
    return tasks, nil
}

func CompleteTask(id int) error {
    tasks, err := ListTasks()
    if err != nil {
        return err
    }

    updated := false
    for i, t := range tasks {
        if t.ID == id {
            tasks[i].Completed = true
            updated = true
        }
    }
    if !updated {
        return fmt.Errorf("no task found with ID %d", id)
    }

    return writeTasks(tasks)
}

func DeleteTask(id int) error {
    tasks, err := ListTasks()
    if err != nil {
        return err
    }

    var updatedTasks []task.Task
    for _, t := range tasks {
        if t.ID != id {
            updatedTasks = append(updatedTasks, t)
        }
    }

    return writeTasks(updatedTasks)
}

func writeTasks(tasks []task.Task) error {
    file, err := os.Create(fileName)
    if err != nil {
        return fmt.Errorf("failed to open file for rewriting: %v", err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, t := range tasks {
        record := []string{strconv.Itoa(t.ID), t.Description, t.Priority, strconv.FormatBool(t.Completed)}
        if err := writer.Write(record); err != nil {
            return fmt.Errorf("failed to write updated record: %v", err)
        }
    }
    return nil
}

func generateRandomID() (int, error) {
    var number int32
    err := binary.Read(rand.Reader, binary.BigEndian, &number)
    if err != nil {
        return 0, fmt.Errorf("failed to generate random number: %v", err)
    }
    return int(number & 0x7FFFFFFF), nil
}
