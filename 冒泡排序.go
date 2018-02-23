package main

import (
	"fmt"
)

func main() {
	array := [5]int{2, 1, 5, 4, 3}
	numlen := len(array)

	fmt.Println("数组array排序前：", array)
	for i := 0; i < numlen; i++ {
		for j := 0; j < i; j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	fmt.Println("数组array排序后：", array)
}
