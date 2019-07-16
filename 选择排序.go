package main

import "fmt"

//选择排序 ：首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，
// 然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序
// 列的末尾。以此类推，直到所有元素均排序完毕
func SelectionSort(s []int) []int {
	n := len(s)

	for i := 0; i < n; i++ {
		max := i
		for j := i + 1; j < n; j++ {
			if s[max] < s[j] {
				max = j
			}
		}
		s[i], s[max] = s[max], s[i]
	}

	return s
}

func main02() {

	s := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	fmt.Println(s)
	//InsertionSort(s)  插入排序
	SelectionSort(s)
	fmt.Println(s)
}

