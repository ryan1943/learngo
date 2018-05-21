package main

import (
	"fmt"
)

func main() {
	arr := []int{-1, -1, 2, 2, 4, 4, 7, 7}
	length := len(arr)
	m := BinFirst(arr, 0, length-1, 4) //第一次出现的下标
	fmt.Println(m)
	n := BinLast(arr, 0, length-1, 4) //最后一次出现的下标
	fmt.Println(n)
}

//非递归二分查找
//返回查找到的位置,-1表示找不到或错误
//时间复杂度O(logN)，空间复杂度O(1)
func BinSearch(arr []int, low, high, k int) int {
	if low < 0 || high < 0 {
		return -1
	}
	for low <= high {
		mid := low + (high-low)>>1
		if k < arr[mid] {
			high = mid - 1
		} else if k > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

//递归二分查找
//时间复杂度O(logN)，空间复杂度O(logN)
func BinSearch2(arr []int, low, high, k int) int {
	if low < 0 || high < 0 {
		return -1
	}
	for low <= high {
		mid := low + (high-low)>>1
		if k < arr[mid] {
			return BinSearch2(arr, low, mid-1, k)
		} else if k > arr[mid] {
			return BinSearch2(arr, mid+1, high, k)
		} else {
			return mid
		}
	}
	return -1
}

//二分查找返回k第一次出现的下标
func BinFirst(arr []int, low, high, k int) int {
	if low < 0 || high < 0 {
		return -1
	}
	for low < high {
		mid := low + (high-low)>>1
		if k > arr[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if arr[low] == k {
		return low
	}
	return -1
}

//二分查找返回k最后一次出现的下标
func BinLast(arr []int, low, high, k int) int {
	if low < 0 || high < 0 {
		return -1
	}
	for low+1 < high {
		mid := low + (high-low)>>1
		if k >= arr[mid] {
			low = mid
		} else {
			high = mid - 1
		}
	}
	if arr[high] == k {
		return high
	} else if arr[low] == k {
		return low
	} else {
		return -1
	}
}
