
package main

import "fmt"

func reverseSlice(s *[]int) {
	start := 0
	end := len(*s)-1
	while start < end {
		temp := (*s)[start]
		(*s)[start] = (*s)[end]
		(*s)[end] = temp
		start += 1
		end -= 1
	}
}

func main() {
	nums := []int{1,2,3,4,5}
	revreseSlice(&nums)
	fmt.Println(nums)
}
