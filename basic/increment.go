
package main

import "fmt"


func increment(x *int){
	*x++
}

func main() {
	x := 7
	increment(&x)
	fmt.Println(x)
}
