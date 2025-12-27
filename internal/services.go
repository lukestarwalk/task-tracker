package internal

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func List(tasks *[]Task) {
	
	for _, task := range *tasks {
		PrintTask(&task)
	}
}

func Add(tasks *[]Task, description string) {
	task := Task{
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format("02-01-2006"),
	}
	task.ID = len(*tasks) + 1;
	*tasks = append(*tasks, task);
	SaveTasks(tasks);
	log.Println("Task created sucssesfully!");
}

func Update(tasks *[]Task, taskId, description string) {

	id, err := strconv.Atoi(taskId);

	if err != nil {
		log.Fatal("Invalid ID!");
	}

	if id > len(*tasks) {
		log.Fatal("There is no task with such id");
	}

	(*tasks)[id-1].Description = description;
	(*tasks)[id-1].UpdatedAt = time.Now().Format("02-01-2006");

	SaveTasks(tasks);
}

func Delete(tasks *[]Task, taskId string) {

	id, err := strconv.Atoi(taskId);

	if err != nil {
		log.Fatal("Invalid ID!");
	}

	for i := id; i < len(*tasks); i++ {
		(*tasks)[i].ID--;
	}

	*tasks = append((*tasks)[:id-1], (*tasks)[id:]...);
	SaveTasks(tasks);
}

func Mark(tasks *[]Task, taskId, typeMark string) {

	id, err := strconv.Atoi(taskId);

	if err != nil {
		log.Fatal("ID Inválido!");
	}

	status := strings.SplitAfter(typeMark,"mark-");

	(*tasks)[id-1].Status = status[1];

	SaveTasks(tasks);
}

func PrintTask(task *Task) {

	fmt.Println("......................................................");
	fmt.Printf("\tID: %d\n\tDescription: %s\n\tSatus: %s\n\tCreated At: %s\n\tUpdated At: %s\n", 
	task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt);
}