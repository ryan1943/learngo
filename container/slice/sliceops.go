package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)
	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)
	fmt.Println("Deleteing elements from slice")
	s2 = append(s2[:3], s2[4:]...)
}
