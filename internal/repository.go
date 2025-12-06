package internal

import (
	"encoding/json"
	"log"
	"os"
)

func save(task Task) {


	jsonData, err := json.MarshalIndent(task, "", "   ");
	
	if err != nil {
		log.Fatal(err);
	}

	file, err := os.Create("tasks");

	if err != nil {
		log.Fatal(err);
	}

	defer file.Close()

	_, err = file.Write(jsonData);

	// file.WriteAt()
	if err != nil {
		log.Fatal(err);
	}

	log.Println("Task created sucssesfully");
}