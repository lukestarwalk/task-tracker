package internal

import (
	"encoding/json"
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
        log.Fatal("Something went wrong while trying to create tasks file:", err)
    }
}

func ReadTasks() *[]Task{

	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
        SaveTasks(nil);
        return nil;
    }

    
    content, err := os.ReadFile("tasks.json");
    if err != nil {
        log.Fatal("S:", err);
    }
    
    if len(content) == 0 {
        return nil;
    }

    var tasks []Task;
    err = json.Unmarshal(content, &tasks);
    if err != nil {
        log.Fatalf("Something went wrong while trying to unmarshal tasks file: %v", err)
    }

    return &tasks;
}


