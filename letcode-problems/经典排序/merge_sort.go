/*
归并排序: 
时间复杂度: o(nlogn)
空间复杂度: o(n)
*/
package main
import "fmt"
// [4, 2, 6, 8, 9, 1]

// merge: [1, 6] <=> [2, 4, 8, 9]

func merge(arr []int, left int, right int, mid int){
	// make slice to copy arr item
	var tmp = make([]int, right - left + 1)
	less := left
	more := mid + 1
	// 临时temp的指针
	t := 0
	// 当都不越界时
	for less <= mid && more <= right {
		if(arr[less] <= arr[more]){
			tmp[t] = arr[less]
			t++
			less++
		}else {
			tmp[t] = arr[more]
			t++
			more++
		}
	}
	// 如果右侧越界，左侧有剩余
	for less <= mid {
		tmp[t] = arr[less]
		t++
		less++
	}
	// 如果右侧有剩余
	for more <= right {
		tmp[t] = arr[more]
		t++
		more++
	}

	// 在left - right范围上copy file
	t = 0
	for left <= right{
		arr[left] = tmp[t]
		t++
		left++
	}
}

func mergeSort(arr []int, left int, right int){
	if left < right {
		mid := left + (right - left) / 2		
		mergeSort(arr, left, mid)
		mergeSort(arr, mid + 1, right)
		merge(arr, left, right, mid)
	}
}

func main(){
	arr := []int{4, 2, 4, 5, 2, 8}
	mergeSort(arr, 0, len(arr) - 1)
	fmt.Println(arr)
}