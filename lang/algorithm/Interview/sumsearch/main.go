package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-5, 8, 7, 3, 1, 0, 2, -1, 4}
	point, ok := SumSearch(arr, 3)
	fmt.Println(point, ok)

	result := SumSearch2(arr, 3)
	fmt.Println(result)

	point, ok = SumSearch3(arr, 3)
	fmt.Println(point, ok)
}

//返回值
type Point struct {
	a, b int
}

//找出数组中的两个数，让这两个数之和等于一个给定的值
//找到第一对数即返回
func SumSearch(arr []int, sum int) (Point, bool) {
	length := len(arr)
	sort.Ints(arr)
	for i := 0; i < length-1; i++ {
		j := BinSearch(arr, i+1, length-1, sum-arr[i])
		if j != -1 {
			return Point{arr[i], arr[j]}, true
		}
	}

	return Point{}, false
}

//找出所有符合的数放到是slice里面一起返回出去
func SumSearch2(arr []int, sum int) []Point {
	var result []Point
	length := len(arr)
	sort.Ints(arr)
	for i := 0; i < length-1; i++ {
		j := BinSearch(arr, i+1, length-1, sum-arr[i])
		if j != -1 {
			result = append(result, Point{arr[i], arr[j]})
		}
	}

	return result
}

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

func SumSearch3(arr []int, sum int) (Point, bool) {
	length := len(arr)
	sort.Ints(arr)
	for i, j := 0, length-1; i < j; {
		if arr[i]+arr[j] == sum {
			return Point{arr[i], arr[j]}, true
		} else if arr[i]+arr[j] < sum {
			i++
		} else {
			j--
		}
	}
	return Point{}, false
}
