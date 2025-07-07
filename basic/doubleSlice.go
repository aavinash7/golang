
package main

import "fmt"


func doubleSlice(x *[]int) {
	for index,_ := range *x {
		(*x)[index] *= 2
	}
}


func main() {
	nums := []int{1,2,3,4}
	doubleSlice(&nums)
	fmt.Println(nums)
}
