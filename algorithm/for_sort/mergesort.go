package main

import "fmt"

// MergeSort 归并排序
func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left []int, right []int) []int {
	fmt.Printf("%v+%v", left, right)
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	fmt.Printf(" -> %v\n", result)
	return result
}

func main() {
	arr := []int{
		22, 34, 3, 32, 82, 55, 89, 50, 37, 5, 64, 35, 9, 70,
	}
	newArr := MergeSort(arr)
	fmt.Println(newArr)
}
