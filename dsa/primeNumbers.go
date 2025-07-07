
package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number: ")
	fmt.Scanln(&num)
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			fmt.Println("Not prime")
			return
		}
	}
	fmt.Println("Prime")

}
