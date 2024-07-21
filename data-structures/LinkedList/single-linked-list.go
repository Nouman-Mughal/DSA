package main

import "fmt"

type Node struct {
	value    interface{}
	next     *Node
	previous *Node // if list is double linked list
}

type LinkedList struct {
	first *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{first: nil}
}

func (ll *LinkedList) AddFirst(value interface{}) {
	node := &Node{value: value}
	node.next = ll.first
	ll.first = node
}

func (ll *LinkedList) AddLast(value interface{}) {
	node := &Node{value: value}
	if ll.first != nil {
		currentNode := ll.first
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = node
	} else {
		ll.first = node
	}
}

func (ll *LinkedList) RemoveLast() *Node {
	var target *Node
	current := ll.first
	if current != nil && current.next != nil {
		for current.next != nil && current.next.next != nil {
			current = current.next
		}
		target = current.next
		current.next = nil
	} else {
		target = ll.first
		ll.first = nil
	}
	return target
}

func (ll *LinkedList) RemoveFirst() interface{} {
	if ll.first != nil {
		first := ll.first
		ll.first = first.next
		return first.value
	}
	return nil
}

func (ll *LinkedList) Contains(value interface{}) int {
	index := 0
	for current := ll.first; current != nil; current = current.next {
		if current.value == value {
			return index
		}
		index++
	}
	return -1
}

func (ll *LinkedList) RemoveAt(nth int) interface{} {
	if nth == 0 {
		return ll.RemoveFirst()
	}
	current := ll.first
	for index := 0; current != nil; index++ {
		if index == nth-1 {
			if current.next == nil {
				return ll.RemoveLast().value
			}
			removedValue := current.next.value
			current.next = current.next.next
			return removedValue
		}
		current = current.next
	}
	return nil
}

func main() {
	// Example usage
	list := NewLinkedList()
	list.AddFirst(1)
	list.AddLast(2)
	list.AddLast(3)
	fmt.Println(list.Contains(2))
	fmt.Println(list.RemoveAt(1))
	fmt.Println(list.RemoveFirst())
}
