package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int) {
	quickSortAlg(arr, 0, len(arr)-1)
}

func quickSortAlg(arr []int, lowIndex, highIndex int) {
	if lowIndex >= highIndex {
		return
	}
	//опорный элемент выбираем случайным образом
	s2 := rand.NewSource(time.Now().Unix())
	r2 := rand.New(s2)
	pivotIndex := r2.Intn(highIndex-lowIndex) + lowIndex
	pivot := arr[pivotIndex]
	arr[pivotIndex], arr[highIndex] = arr[highIndex], arr[pivotIndex]

	leftP := lowIndex
	rightP := highIndex - 1
	for leftP < rightP {
		for arr[leftP] <= pivot && leftP < rightP {
			leftP++
		}
		for arr[rightP] >= pivot && leftP < rightP {
			rightP--
		}
		//свап
		arr[leftP], arr[rightP] = arr[rightP], arr[leftP]
	}
	if arr[leftP] > arr[highIndex] {
		arr[leftP], arr[highIndex] = arr[highIndex], arr[leftP]
	} else {
		leftP = highIndex
	}
	quickSortAlg(arr, lowIndex, leftP-1)
	quickSortAlg(arr, leftP+1, highIndex)

}

func main() {
	s1 := rand.NewSource(time.Now().Unix())
	r1 := rand.New(s1)

	const n = 15
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = r1.Intn(50)
	}
	fmt.Println("----- Before:  -----")
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	quickSort(a)
	fmt.Println("\n----- After:  -----")
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
}
