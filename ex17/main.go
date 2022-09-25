package main

import "fmt"

func binarySearch(arr []int, target int) int {
	return binarySearchAlg(arr, target, 0, len(arr)-1)
}

// рекурсия
func binarySearchAlg(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	} else if target < arr[mid] {
		return binarySearchAlg(arr, target, left, mid-1)
	} else {
		return binarySearchAlg(arr, target, mid+1, right)
	}
}

// итеративно
func binarySearchIter(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if target < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 6, 8, 10, 12, 15, 20}
	fmt.Println(arr)
	fmt.Printf("Index = %d\n", binarySearch(arr, 12))                //индекс = 5
	fmt.Printf("Iterative, index = %d\n", binarySearchIter(arr, 20)) // индекс = 7

}
