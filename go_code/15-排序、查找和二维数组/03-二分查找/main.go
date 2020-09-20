package main

import "fmt"

// 二分查找
func BinarySearch(arr *[6]int, target int) int {
	if (*arr)[0] > target || (*arr)[5] < target {
		return -1
	}
	var (
		left  = 0
		right = len(arr) - 1
	)
	for ; left < right; {
		var middle = (left + right) / 2
		if (*arr)[middle] > target {
			right = middle - 1
		} else if (*arr)[middle] < target {
			left = middle + 1
		} else {
			return middle
		}

	}
	return -1
}

func RecurseBinarySearch(arr *[6]int, target int, left int, right int) int {
	if left > right {
		return -1
	}
	middle := left + (right-left)/2
	if (*arr)[middle] > target {
		right = middle - 1
		return RecurseBinarySearch(arr, target, left, right)
	} else if (*arr)[middle] < target {
		left = middle + 1
		return RecurseBinarySearch(arr, target, left, right)
	} else {
		return middle
	}
}

func main() {
	arr := [...]int{1, 8, 10, 89, 1000, 1234}
	fmt.Println(BinarySearch(&arr, 10))
	fmt.Println(RecurseBinarySearch(&arr, 10, 0, 5))
}
