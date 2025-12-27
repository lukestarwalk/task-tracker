package main

import (
	"log"
	"os"
	"task-tracker/internal"
)

type Task = internal.Task;

func main() {
	
	args := os.Args;
	numArgs := len(args);

	if numArgs < 2 {
		log.Fatal("Necessita de pelo menos dois argumentos!");
	}
	
	option := args[1];

	tasks := internal.ReadTasks();
	switch option {

	case "list":
		internal.List(tasks);
		
	case "add":
		if numArgs != 3 {
			log.Fatal("It's required 3 arguments");
		}
		internal.Add(tasks, args[2]);

	case "update":
		if numArgs != 4 {
			log.Fatal("It's required 4 arguments");
		}
		internal.Update(tasks, args[2], args[3]);

	case "delete":

		if numArgs != 3 {
			log.Fatal("It's required 3 arguments");
		}

		internal.Delete(tasks, args[2]);

	case "mark-in-progress":

		if numArgs != 3 {
			log.Fatal("It's required 3 arguments");
		}
		internal.Mark(tasks, args[2], "mark-in-progress");

	case "mark-done":

		if numArgs != 3 {
			log.Fatal("It's required 3 arguments");
		}
		internal.Mark(tasks, args[2], "mark-done");

	default:
		log.Fatal("Inavlid type of argument!");
	}
}