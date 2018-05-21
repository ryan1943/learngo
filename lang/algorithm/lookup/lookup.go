package main

func main() {

}

func binarySearch(arr []int, start, end, k int) int {
	for start <= end {
		mid := start + (end-start)/2
		if k < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}
