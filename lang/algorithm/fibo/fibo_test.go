package main

import "testing"

var testArr = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

func TestFibo1(t *testing.T) {
	for i, v := range testArr {
		if actual := Fibo1(i); actual != v {
			t.Errorf("Fibo1(%v): got %v, expected %v\n", i, actual, v)
		}
	}
}

func TestFibo2(t *testing.T) {
	for i, v := range testArr {
		if actual := Fibo2(i); actual != v {
			t.Errorf("Fibo2(%v): got %v, expected %v\n", i, actual, v)
		}
	}
}

func TestFibo3(t *testing.T) {
	for i, v := range testArr {
		if actual := Fibo3(i); actual != v {
			t.Errorf("Fibo3(%v): got %v, expected %v\n", i, actual, v)
		}
	}
}
