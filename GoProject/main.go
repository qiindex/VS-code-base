package main

import "fmt"

func main() {
	name := "hello go"
	fmt.Println("go world", name)
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7
	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Printf("目标数字 %d 在数组中的索引位置是: %d\n", target, result)
	} else {
		fmt.Printf("目标数字 %d 不在数组中\n", target)
	}
}
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
	

