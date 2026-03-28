package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task struct {
	id          int
	name        string
	description string
	complete    bool
	date        time.Time
}

type Note struct {
	id          int
	name        string
	description string
	tasks       []*Task
}

func (n *Note) completeTask(task *Task) {
	(*task).complete = true
}

func (n *Note) createTask(name string, description string) (newTask *Task) {
	id := generateId()
	newTask = &Task{id, name, description, false, time.Now()}
	n.tasks = append(n.tasks, newTask)
	return
}

func main() {
	note := &Note{
		id:          generateId(),
		name:        "aprender go",
		description: "quiero aprender go",
		tasks:       []*Task{},
	}

	task := note.createTask("structs", "aprender structs")

	fmt.Println(note.tasks[0])

	note.completeTask(task)

	fmt.Println(note.tasks[0])

}

func generateId() int {
	return rand.Intn(10000) + 1
}
