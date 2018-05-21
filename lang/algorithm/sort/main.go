package main

import "fmt"

func main() {
	arr := []int{8, 4, 2, 9, 10, -3, 0, 20, 15, -1}
	//BubbleSort(arr)
	//SelectSort(arr)
	//InsertSort(arr)
	//ShellSort(arr)
	//HeapSort(arr)
	//MergeSort(arr)
	length := len(arr)
	QuickSort2(arr, 0, length-1)
	fmt.Println(arr)
}

//冒泡排序
//对相邻的元素进行两两比较，顺序相反则进行交换
//每一趟把最大的元素放在末尾
//时间复杂度O(n^2)
func BubbleSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		flag := true //若为true，则表示此次循环没有进行交换，也就是待排序列已经有序
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

//简单选择排序
//时间复杂度O(n^2)
func SelectSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i //每一趟的开始把首元素的下标作为最小元素的下标
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}

//插入排序
//把第一个元素作为有序序列
//每一趟将一个待排序的数，通过从后往前两两比较
//插入到前面已经排好序的有序序列中去
//最好的情况时间复杂度为O(n);在最坏情况下，时间复杂度为O(n^2)
func InsertSort(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		j := i
		if arr[i] < arr[i-1] {
			for j > 0 && arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				j--
			}
		}
	}
}

//希尔排序
//按下标的一定增量分组，对每组使用直接插入排序算法排序
//平均时间复杂度O(nlog2n)
//最坏时间复杂度依然为O(n2)
func ShellSort(arr []int) {
	length := len(arr)
	//增量gap，并逐步缩小增量
	for gap := length / 2; gap > 0; gap /= 2 {
		//从第gap个元素，逐个对其所在组进行直接插入排序操作
		for i := gap; i < length; i++ {
			j := i
			for j-gap >= 0 && arr[j] < arr[j-gap] {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
				j -= gap
			}
		}
	}
}

//堆排序
//时间复杂度O(log2n)
func HeapSort(a []int) {
	length := len(a)
	if length == 0 {
		return
	}
	//构造初始堆
	for i := length/2 - 1; i >= 0; i-- {
		heapAdjust(a, i, length-1)
	}

	for j := length - 1; j >= 0; j-- {
		a[0], a[j] = a[j], a[0]
		heapAdjust(a, 0, j-1)
	}
}

//调整堆
func heapAdjust(a []int, start, end int) {
	temp := a[start]

	for k := 2*start + 1; k <= end; k = 2*k + 1 { //从i结点的左子结点开始，也就是2i+1处开始
		//选择出左右孩子较大的下标
		if k < end && a[k] < a[k+1] {
			k++
		}
		//如果子节点大于父节点，将子节点值赋给父节点（不用进行交换）
		if a[k] > temp {
			a[start] = a[k]
			start = k
		} else {
			break
		}
	}
	a[start] = temp //插入正确的位置
}

//归并排序
func MergeSort(arr []int) {
	length := len(arr)
	temp := make([]int, length)
	mSort(arr, 0, length-1, temp)
}

func mSort(arr []int, left, right int, temp []int) {
	if left < right {
		mid := (left + right) / 2
		mSort(arr, left, mid, temp)
		mSort(arr, mid+1, right, temp)
		//两边的子序列都是有序的，
		//如果左边的最大的元素比右边最小的元素大才需要合并
		if arr[mid] > arr[mid+1] {
			merge(arr, left, mid, right, temp)
		}
	}
}

func merge(arr []int, left, mid, right int, temp []int) {
	i := left
	j := mid + 1
	t := 0 //临时slice的指针
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[t] = arr[i]
			i++
		} else {
			temp[t] = arr[j]
			j++
		}
		t++
	}
	//将左序列剩余元素填充进temp中
	for i <= mid {
		temp[t] = arr[i]
		t++
		i++
	}
	//将右序列剩余元素填充进temp中
	for j <= right {
		temp[t] = arr[j]
		t++
		j++
	}
	t = 0
	//将temp中的元素全部拷贝到原数组中
	for left <= right {
		arr[left] = temp[t]
		left++
		t++
	}
}

//快速排序
func QuickSort1(arr []int, left, right int) {
	if left < right {
		i := arrAdjust(arr, left, right)
		QuickSort1(arr, left, i-1)
		QuickSort1(arr, i+1, right)
	}
}

//返回调整后基准数的位置
func arrAdjust(arr []int, left, right int) int {
	i, j := left, right
	x := arr[left] //基准
	for i < j {
		//从右向左找小于x的数来填arr[i]
		for i < j && arr[j] >= x {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}

		//从左向右找大于或等于x的数来填arr[j]
		for i < j && arr[i] < x {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = x //退出时，i等于j,将x填到这里

	return i
}

//改进版的快速排序
func QuickSort2(arr []int, left, right int) {
	mid := (left + right) / 2
	arr[left], arr[mid] = arr[mid], arr[left] //可以选择中间的数作为基准
	x := arr[left]                            //基准
	i, j := left, right
	for i < j {
		for i < j {
			//从右向左找小于x的数来填arr[i]
			for i < j && arr[j] >= x {
				j--
			}
			if i < j {
				arr[i] = arr[j]
				i++
			}

			//从左向右找大于或等于x的数来填arr[j]
			for i < j && arr[i] < x {
				i++
			}
			if i < j {
				arr[j] = arr[i]
				j--
			}
		}
		arr[i] = x
		QuickSort2(arr, left, i-1)
		QuickSort2(arr, i+1, right)
	}
}
