package main

import "fmt"

func BinarySearch(array []int, target int) int {
	n := len(array)
	low, high := 0, n-1
	for low <= high {
		mid := low + (high-low)/2
		switch {
		case array[mid] < target:
			low = mid + 1
		case array[mid] > target:
			high = mid - 1
		default:
			return mid
		}
	}
	return -1
}

func main() {
	fmt.Println(BinarySearch([]int{1, 3, 5, 6, 7}, 6))
	fmt.Println(BinarySearch([]int{1, 3, 5, 6, 7}, 7))
	fmt.Println(BinarySearch([]int{1, 3, 5, 6, 7}, 9))
	fmt.Println(BinarySearch([]int{1, 3, 5, 6, 7}, 1))
	fmt.Println(BinarySearch([]int{1, 3, 5, 6, 7}, 0))
}
