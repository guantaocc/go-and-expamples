/*
冒泡排序: 从低到高
*/
package main

func bubbleSort(arr []int) {
	for i:= len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if(arr[j+1] > arr[j]){
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		} 
	}
}