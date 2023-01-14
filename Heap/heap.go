//insert and remove are O(h) where h is the hight of the tree
//to represetn the time complexity in N as an array, its O(logN) since the hight and the number of node has a logarithmic relation

//swapping the nodes to the max/min heap format is called heapifying

package main

import "fmt"

//MaxHeap struct is made of a slice that holds an array
type MaxHeap struct {
	array []int
}

//insert adds an element to the heap
func (h *MaxHeap) insert(key int) {
	//will add elem to the array then heapify
	h.array = append(h.array, key)
	//heapify the last elem
	h.maxHeapifyUp(len(h.array) - 1)
}

//remove returns the largest key (root), and removes it from the heap
func (h *MaxHeap) remove() int {
	toRemove := h.array[0]
	size := len(h.array) - 1

	//if array is empty
	if len(h.array) == 0 {
		fmt.Println("cannot remove because heap is empty")
		return -1
	}

	//replace the right most leaf child as the new root
	h.array[0] = h.array[size]
	h.array = h.array[:size] //resize array

	h.maxHeapifyDown(0) //from the 0th index

	return toRemove

}

//maxHeapifyingDown will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	left, right := getLeftChild(index), getRightChild(index)
	childToCompare := 0

	//loop whileindex has at least one child
	for left <= lastIndex {
		if left == lastIndex { //when the left child is the ony child
			childToCompare = left
		} else if h.array[left] > h.array[right] { //when the left child is larger
			childToCompare = left
		} else { //when the right child is larger
			childToCompare = right
		}
		//compare value of current index  to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			left, right = getLeftChild(index), getRightChild(index)
		} else { //index has found correct place, stop heapifying
			return
		}
	}

}

//maxHeapifyingUp will heapify from buttom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[getParent(index)] < h.array[index] {
		h.swap(getParent(index), index)
		//update the index to the parent
		index = getParent(index)
	}

}

//swap keys in array
func (h *MaxHeap) swap(parentIndex int, childIndex int) {
	h.array[parentIndex], h.array[childIndex] = h.array[childIndex], h.array[parentIndex]
}

//get parent index from child index
//rounding down will yeild the same parent index
//leftchild=2n+1 rightchild=2n+2 where n is the index of the parent index
func getParent(childIndex int) int {
	//(leftChild - 1)/2 = parentIndex
	return (childIndex - 1) / 2
}

//leftCHild is always an odd num (+1)
func getLeftChild(parentIndex int) int {
	return 2*parentIndex + 1
}

//rightChild is always an even num (+2)
func getRightChild(parentIndex int) int {
	return 2*parentIndex + 2
}

func main() {
	heapBuilder := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	myHeap := &MaxHeap{}
	for _, value := range heapBuilder {
		myHeap.insert(value)
		fmt.Println(myHeap)

	}
	//fmt.Println(myHeap)
	fmt.Println()

	for i := 0; i < 5; i++ {
		myHeap.remove()
		fmt.Println(myHeap)

	}

}
