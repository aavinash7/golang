
package main

import "fmt"

func swap(a,b *int){
	c := *a
	*a = *b
	*b = c
	fmt.Println(*a)
	fmt.Println(*b)
}

func main() {
	a := 5
	b := 10
	swap(&a, &b)
	fmt.Println(a,b)
}
