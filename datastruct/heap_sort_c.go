package main

import "fmt"

// 堆操作

// 堆的二叉树高度 o(logN)时间复杂度

// 大根堆插入 heapInsert过程
func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

// heapfiy过程
func heapify(arr []int, index int, heapSize int) {
	left := index*2 + 1
	for left < heapSize {
		// 当有子孩子时候, 取较大值的下标
		var largest int
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		} else {
			largest = left
		}
		// 比较和父节点大小
		if arr[largest] < arr[index] {
			largest = index
		}
		// 如果都不比父亲大
		if largest == index {
			break
		}
		// 交换值
		arr[largest], arr[index] = arr[index], arr[largest]
		// 继续判断
		index = largest
		// 左节点
		left = index*2 + 1
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 堆排序
func heapSort(arr []int) {
	if arr == nil || len(arr) <= 2 {
		return
	}
	for i := len(arr) - 1; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}
	heapSize := len(arr)
	heapSize--
	swap(arr, 0, heapSize)
	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		heapSize--
		swap(arr, 0, heapSize)
	}
}

func main() {
	arr := []int{5, 6, 1, 4, 3, 2}
	heapSort(arr)
	fmt.Print(arr)
}
