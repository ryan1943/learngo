package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if r == nil {
			fmt.Println("Nothing to recover. Please try uncomment errors below.")
			return
		}
		if err, ok := r.(error); ok {
			fmt.Println("Error occured: ", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()
	//newErr := errors.New("hahahah")
	//panic(newErr)
	//panic(123)
}

func main() {
	tryRecover()
}
