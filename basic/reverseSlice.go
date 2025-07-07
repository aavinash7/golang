
package main

import "fmt"


func reverseSlice(s *[]int){
	start := 0
	end := len(*s)-1
	for start < end {
		temp := (*s)[start]
		(*s)[start] = (*s)[end]
		(*s)[end] = temp
		start++
		end--
	}
}

func main() {
	nums := []int{1,2,3,4,5}
	reverseSlice(&nums)
	fmt.Println(nums)
}
