package linkedlist

import (
	"errors"
	"fmt"
)

type ElemType interface{}

//节点
type Node struct {
	Data ElemType
	Next *Node
}

//链表
type LinkedList struct {
	Head *Node //头结点
}

func CreateLinkedList() *LinkedList {
	head := new(Node)
	return &LinkedList{head}
}

//获取第i个结点的值
func (list *LinkedList) GetELem(i int) (ElemType, bool) {
	j := 1
	p := list.Head.Next     //第一个结点
	for p != nil && j < i { //p不等于nil或者j不等于i时，循环继续
		p = p.Next
		j++
	}
	if p == nil || j > i {
		return nil, false //第i个结点不存在
	}

	return p.Data, true
}

//在第i个位置前插入元素
func (list *LinkedList) Insert(i int, x ElemType) bool {
	j := 1
	p := list.Head          //第0个结点
	for p != nil && j < i { //寻找第i-1个结点
		p = p.Next
		j++
	}
	if p == nil || j > i {
		return false //第i个结点不存在
	}
	s := &Node{x, p.Next}
	p.Next = s

	return true
}

//删除第i个结点
func (list *LinkedList) Delete(i int) (ElemType, bool) {
	j := 1
	p := list.Head
	for p.Next != nil && j < i { //寻找第i-1个结点
		p = p.Next
		j++
	}
	if p.Next == nil || j > i {
		return nil, false //第i个结点不存在
	}
	q := p.Next
	p.Next = q.Next

	return p.Data, true
}

//判断链表是否为空
func (list *LinkedList) IsEmpty() bool {
	return list.Head.Next == nil
}

//链表长度
func (list *LinkedList) Len() int {
	length := 0
	p := list.Head.Next
	for p != nil {
		p = p.Next
		length++
	}

	return length
}

//打印链表
func (list *LinkedList) Print() error {
	if list.IsEmpty() {
		return errors.New("this is an empty linked list")
	}
	p := list.Head.Next
	for p != nil {
		fmt.Printf("%v  ", p.Data)
		p = p.Next
	}
	fmt.Println()
	return nil
}
