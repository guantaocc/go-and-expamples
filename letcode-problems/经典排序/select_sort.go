/*
选择排序：
选择一个数到应有位置上
时间复杂度: o(n2)
空间: o(1)
*/
package main

func selectSort(arr []int){
	for i:= 0; i < len(arr); i++ {
		for j:= i; j < len(arr); j++ {
			if(arr[j] < arr[i]){
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}