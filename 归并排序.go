package main

import "fmt"

//归并排序

//分组
func mergeSort(arr []int) []int {
	//判断退出条件
	if len(arr) < 2 {
		return arr
	}
	//获取排序切片长度
	n := len(arr) / 2

	//继续分组
	left := mergeSort(arr[0:n])
	right := mergeSort(arr[n:])

	//合并
	result := merge(left, right)

	return result

}

func merge(left, right []int) []int {
	//定义合并容器
	result := make([]int, 0)

	//下标值
	m, n := 0, 0

	for m < len(left) && n < len(right) {
		if left[m] < right[n] {
			result = append(result, left[m])
			m++
		} else {
			result = append(result, right[n])
			n++
		}
	} //for

	result = append(result, left[m:]...)
	result = append(result, right[n:]...)

	return result
}

func main() {
	arr := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	fmt.Println(arr)
	result := mergeSort(arr)
	fmt.Println(result)
}
