package main

import "fmt"

type LinkList struct {
	Val interface{}
	Next *LinkList
}

func main() {
	list := LinkList{Val:nil}
	list.New(2)
	list.New(3)
	last := list.GetLastNode()
	fmt.Println(last.Val)
}

func (list *LinkList) New(val interface{})  {
	new := LinkList{Val:val}
	last := list.GetLastNode()
	last.Next = &new
}

func (list *LinkList) GetLastNode() *LinkList {
	last := list
	for  {
		if last.Next!=nil {
			last = last.Next
		}else {
			return last
		}
	}
}