package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Name  string
	Date  string
	State bool
}

func AddTasks() {
	var (
		taskName string
		taskDate string
	)
	fmt.Print("Task name : ")
	fmt.Scanln(&taskName)
	fmt.Print("Task Date : ")
	fmt.Scanln(&taskDate)

	userTask := Task{
		Name:  taskName,
		Date:  taskDate,
		State: false,
	}

	data, err := json.Marshal(userTask)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
	// creating a task.json file for storing the tasks
	newdata , err := os.Stat("./tasks.json") ; os.IsNotExist(err)

	fmt.Println(newdata)
	// os.WriteFile("tasks.json", []byte(data), 0644)
}

func EditTasks() {

}

func DeleteTasks() {

}
