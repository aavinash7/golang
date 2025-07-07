
package main

import "fmt"

func rotateLeft(s *[]int) {
	start := (*s)[0]
	for i := 0; i < len(*s)-1; i++ {
		(*s)[i] = (*s)[i+1] 
	}
	(*s)[len(*s)-1] = start
}

func main() {
	nums := []int{1,2,3,4,5}
	rotateLeft(&nums)
	fmt.Println(nums)
}
