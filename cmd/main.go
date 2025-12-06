package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task-tracker/internal"
)

var i rune;

type Task = internal.Task;

func main() {
	
	tasks := read();
	fmt.Println(tasks);

	// task := internal.Task{
	// 	ID: 2,
	// 	Description: "Estudar Go",
	// 	Status: "todo",
	// 	CreatedAt: "01-10-2025",
	// 	UpdatedAt: "25-12-2025",
	// }

	// save(task);

	
	// command := args[1];

	// switch command {
	// case "add":
	// case "update":
	// case "delete":
	// case "mark-in-progress":
	// case "mark-done":
	// default:
	// 	log.Fatal("Inavlid type of argument!");
	// }
}

func read() []Task{

	content, err := os.ReadFile("tasks");

	if err != nil {
		log.Fatal(err);
	}

	var tasks []Task;

	err = json.Unmarshal(content, &tasks)
    if err != nil {
        log.Fatalf("Erro ao fazer Unmarshal do JSON: %v", err)
    }

	fmt.Println(tasks);

	return tasks;
}

func save(task Task) {

	jsonData, err := json.MarshalIndent(task, "", "   ");
	
	if err != nil {
		log.Fatal(err);
	}

	file, err := os.Create("tasks.json");

	if err != nil {
		log.Fatal(err);
	}

	defer file.Close()

	_, err = file.Write(jsonData);

	if err != nil {
		log.Fatal(err);
	}

	log.Println("Task created sucssesfully!");
}