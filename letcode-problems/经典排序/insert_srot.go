/*
插入排序:
将每个位置上数插入到应该在的位置上
时间复杂度o(n2)
*/
package main

func insertSort(arr []int){
	for i:= 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j + 1] < arr[j]; j--{
			arr[j+1], arr[j] = arr[j], arr[j+1]
		}
	}
}