package main

import "fmt"

func rotateRight(s *[]int){
	end := (*s)[len(*s)-1]
	for i := len(*s)-1; i > 0; i-- {
		(*s)[i] = (*s)[i-1]
	}		
	(*s)[0] = end
}

func main() {
	nums := []int{1,2,3,4,5}
	rotateRight(&nums)
	fmt.Println(nums)
}
