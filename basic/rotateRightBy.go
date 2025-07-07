
package main

import "fmt"

func rotateRightBy(n int,s *[]int) {
	n = n % len(*s)
	for range n {
		temp := (*s)[len(*s)-1]
		for i := len(*s)-1; i > 0 ; i-- {
			(*s)[i] = (*s)[i-1]
		}
		(*s)[0] = temp
	}
}


func main() {
	nums := []int{1,2,3,4,5}
	rotate := 13
	rotateRightBy(rotate,&nums)
	fmt.Println(nums)
}



// 1,2,3,4,5
// 5,1,2,3,4
// 4,5,1,2,3
// 3,4,5,1,2
