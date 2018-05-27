package main

import (
	"learngo/lang/datastruct/doublelinkedlist"
)

func main() {
	list := doubleLinkedList.CreateList()
	s := []int{1, 2, 3, 4, 5, 6, 7}
	for _, v := range s {
		list.PushBack(v)
	}
	//list.Print()
	list.Modify(5, 55)
	list.Print()
	list.DeleteVal(7)
	list.Print()
	list.PopBack()
	list.Print()
	list.PopFront()
	list.Print()
	list.Reverse()
	list.Print()

}
