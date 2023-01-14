package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

// without pointer, we would be altering a copy of the LL object
func (l *LinkedList) prepend(newNode *node) {
	//make original head the second node in list
	second := l.head
	//set newNode to head
	l.head = newNode
	//point the newNode to the rest of the list
	l.head.next = second
	l.length++
}

func (l LinkedList) printList() {
	nodeToPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d -> ", nodeToPrint.data)

		nodeToPrint = nodeToPrint.next
		l.length--
	}
	fmt.Println()
}

func (l *LinkedList) delete(value int) {

	//case for empty list
	if l.length == 0 {
		fmt.Println("linked list is empty")
		return
	}

	//case for delete head
	if l.head.data == value {
		//make second node in list the head
		l.head = l.head.next
		l.length--
		fmt.Println("head of list is deleted")
		return
	}

	curr := l.head
	//must traverse till we find the value of the node aheaf of the current node since a singly-linked list cannot move backwards
	for curr.next.data != value {
		//case for value not in list
		if curr.next.next == nil {
			fmt.Println("value not in list")
			return
		}
		curr = curr.next
	}
	//reassgin by skipping the node ahead of the current node
	curr.next = curr.next.next
	l.length--
	fmt.Println("sucessfull deletion")
}

func main() {
	myList := LinkedList{}
	//myList2 := linkedList{}
	node1 := &node{data: 1}
	node2 := &node{data: 6}
	node3 := &node{data: 19}
	node4 := &node{data: 7}
	node5 := &node{data: 37}
	node6 := &node{data: 24}
	node7 := &node{data: 16}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.prepend(node7)

	//fmt.Println(myList)
	myList.printList()
	myList.delete(234567)
	myList.printList()

}
