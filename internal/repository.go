package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func SaveTasks(tasks *[]Task) {

	jsonData, err := json.MarshalIndent(tasks, "", "   ");
	
	if err != nil {
		log.Fatal(err);
	}

	err = os.WriteFile("tasks.json", jsonData, 0644)
    if err != nil {
        log.Fatal("Erro ao escrever arquivo:", err)
    }
}

func ReadTasks() *[]Task{

	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
        SaveTasks(nil);
        return nil;
    }

    
    content, err := os.ReadFile("tasks.json");
    if err != nil {
        log.Fatal("Erro ao ler arquivo:", err);
    }
    
    if len(content) == 0 {
        return nil;
    }

    var tasks []Task;
    err = json.Unmarshal(content, tasks);
    if err != nil {
        log.Fatalf("Erro ao fazer Unmarshal do JSON aqui: %v", err)
    }

    return &tasks;
}

func PrintTask(task *Task) {

	fmt.Println("......................................................");
	fmt.Printf("\tID: %d\n\tDescription: %s\n\tSatus: %s\n\tCreated At: %s\n\tUpdated At: %s\n", 
	task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt);
}

