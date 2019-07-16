//package main
//
//import "fmt"
//
//func QuickSort(nums []int, start, end int) []int{
//	if start >= end {
//		return nil
//	}
//
//	mid := partition(nums, start, end)
//	QuickSort(nums, start, mid)
//	QuickSort(nums, mid+1, end)
//
//	return nums
//}
//
//func partition(nums []int, start, end int) int {
//	temp := nums[start]
//
//	low := start
//	height := end
//
//	for low < height {
//		for low < height && temp < nums[height] {
//			height--
//		} //for
//
//		if low < height {
//			nums[low] = nums[height]
//		} //if
//
//		for low < height && temp > nums[low] {
//			low++
//		} //for
//
//		if low < height {
//			nums[height] = nums[low]
//		} //if
//
//	} //for
//
//	nums[low] = temp
//
//	return low
//}
//
//func main() {
//	arr := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
//	fmt.Println(arr)
//	arr = QuickSort(arr, 0, len(arr)-1)
//	fmt.Println(arr)
//}

package main

import "math/rand"
import "fmt"
import "time"

func quickSort(values []int, left, right int) {

	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && values[i] <= temp {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}

	}

	values[p] = temp

	if p-left > 1 {
		quickSort(values, left, p-1)
		//go quickSort(values,left,p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
		//go quickSort(values,p+1,right)
	}

}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}

func main() {
	ceshi := make([]int, 100)
	for i := 0; i < 100; i++ {
		ceshi[i] = rand.Intn(100)
	}
	start_time := time.Now()
	QuickSort(ceshi)
	end_time := time.Since(start_time)
	fmt.Println("after sort", ceshi)
	fmt.Println("count time ", end_time)
}
