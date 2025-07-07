package main

import "fmt"

func main() {
	type Task struct {
		Name      string
		Completed bool
	}
	tasks := []Task{{"Wakeup", true}, {"Eat", true}, {"Sleep", false}}
	for {
		fmt.Println("Enter any of these --> add or list or complete or exit")
		var user_input string
		var task string
		var complete_task int
		fmt.Scanln(&user_input)
		if user_input == "add" {
			fmt.Println("Enter the task: ")
			fmt.Scanln(&task)
			tasks = append(tasks, Task{task, false})
		} else if user_input == "list" {
			for index, task := range tasks {
				if task.Completed {
					fmt.Printf("%d. [X] %s\n", index+1, task.Name)
				} else {
					fmt.Printf("%d. [ ] %s\n", index+1, task.Name)
				}
			}
		} else if user_input == "complete" {
			fmt.Println("Which task number to be complete?")
			fmt.Scanln(&complete_task)
			if complete_task > 0 && complete_task <= len(tasks) {
				tasks[complete_task-1].Completed = true
			} else {
				fmt.Println("Invalid task number.")
			}
		} else {
			break
		}
	}
}
