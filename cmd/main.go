package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"task-tracker/internal"
	"time"
	// "time"
)

type Task = internal.Task;

func main() {
	
	args := os.Args;
	numArgs := len(args);

	if numArgs < 2 {
		log.Fatal("Necessita de pelo menos dois argumentos!");
	}
	
	option := args[1];
	fmt.Println(option);
	switch option {

	case "list":

		tasks := ReadTasks();

		for _, task := range tasks {
			PrintTask(task);
		}


		
	case "add":
		
		if numArgs < 3 || numArgs > 3 {
			log.Fatal("Para adicionar só necessita de 3 argumentos");
		}
		
		task := Task{
			Description: args[2],
			Status: "todo",
			CreatedAt: fmt.Sprint(time.Now().Format("02/01/2006")),
		}
		AddTask(task);

	case "update":
	case "delete":

		if numArgs < 3 || numArgs > 3 {
			log.Fatal("Para adicionar só necessita de 3 argumentos");
		}

		id, err := strconv.Atoi(args[2]);

		if err != nil {
			log.Fatal("ID Inválido!");
		}
		tasks := ReadTasks();
		for i := id; i < len(tasks); i++ {
			tasks[i].ID--;
		}

		tasks = append(tasks[:id-1], tasks[id:]...);
		SaveTask(tasks);

	case "mark-in-progress":

		if numArgs < 3 || numArgs > 3 {
			log.Fatal("Para adicionar só necessita de 3 argumentos");
		}

		id, err := strconv.Atoi(args[2]);

		if err != nil {
			log.Fatal("ID Inválido!");
		}

		tasks := ReadTasks();
		tasks[id-1].Status = "in-progress";
		SaveTask(tasks);

	case "mark-done":

		if numArgs < 3 || numArgs > 3 {
			log.Fatal("Para adicionar só necessita de 3 argumentos");
		}

		id, err := strconv.Atoi(args[2]);

		if err != nil {
			log.Fatal("ID Inválido!");
		}

		tasks := ReadTasks();
		tasks[id-1].Status = "done";
		SaveTask(tasks);

	default:
		log.Fatal("Inavlid type of argument!");
	}
}

func ReadTasks() []Task{

	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
        SaveTask([]Task{});
        return []Task{}
    }

    content, err := os.ReadFile("tasks.json");
    if err != nil {
        log.Fatal("Erro ao ler arquivo:", err);
    }

    if len(content) == 0 {
        return []Task{};
    }

    var tasks []Task;
    err = json.Unmarshal(content, &tasks);
    if err != nil {
        log.Fatalf("Erro ao fazer Unmarshal do JSON: %v", err)
    }

    return tasks
}

func SaveTask(tasks []Task) {

	jsonData, err := json.MarshalIndent(tasks, "", "   ");
	
	if err != nil {
		log.Fatal(err);
	}

	err = os.WriteFile("tasks.json", jsonData, 0644)
    if err != nil {
        log.Fatal("Erro ao escrever arquivo:", err)
    }

}

func AddTask(task Task) {
	tasks := ReadTasks();
	task.ID = len(tasks) + 1;
	tasks = append(tasks, task);
	SaveTask(tasks);
	log.Println("Task created sucssesfully!");
}

func PrintTask(task Task) {

	fmt.Println("......................................................");
	fmt.Printf("\tID: %d\n\tDescription: %s\n\tSatus: %s\n\tCreated At: %s\n\tUpdated At: %s\n", 
	task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt);
	fmt.Println("......................................................");
}