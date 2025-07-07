
package main

import "fmt"

func main(){
	secret_num := 7
	var num int
	res := false
	fmt.Println("Enter your number: ")
	for i := 0; i < 3; i++ {
		fmt.Scanln(&num)
		if num == secret_num {
			fmt.Println("You win!")
			res = true
			break
		} else {
			if num > secret_num {
				fmt.Println("Too High!")
			} else {
				fmt.Println("Too Low!")
			}
		}
	}
	if !res {
		fmt.Printf("You lose! The number was %d\n",secret_num)
	}
}
