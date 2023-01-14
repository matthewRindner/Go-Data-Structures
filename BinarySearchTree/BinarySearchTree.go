package main

import "fmt"

var count int

//define node class
type Node struct {
	key   int
	left  *Node
	right *Node
}

//insert
//the key should not already be in the tree
func (n *Node) insert(k int) {
	if n.key < k {
		//move right
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.insert(k)
		}
	} else if n.key > k {
		//move left
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.insert(k)
		}
	}
}

//search
func (n *Node) search(k int) bool {
	count++

	//no match
	if n == nil {
		return false
	}
	if n.key < k {
		//move right
		return n.right.search(k)
	} else if n.key > k {
		//move left
		return n.left.search(k)
	}

	return true
}

//remove

func main() {
	//BST with only root node
	myBST := &Node{key: 67}
	myBST.insert(50)
	myBST.insert(33)
	myBST.insert(87)
	myBST.insert(38)
	myBST.insert(23)
	fmt.Println(myBST)
	fmt.Println(myBST.search(2345), count)
}
