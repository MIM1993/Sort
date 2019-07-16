package main

import "fmt"

//冒泡排序
func BubbleSort(array []int) []int {
	n := len(array)
	if n <= 0 {
		return nil
	}

	for i := 0; i < n; i++ {
		//提前推出标志
		flag := false
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				flag = true
			}
		}
		//判断推出
		if !flag {
			break //用于循环语句中跳出循环,并开始执行循环之后的语句
		}
	}

	return array
}

func main() {
	arr := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	BubbleSort(arr)
	fmt.Println(arr)
}
