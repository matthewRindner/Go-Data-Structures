package main

import "fmt"

const Array_Size = 7

// HashTable will hold an array
type HashTable struct {
	array [Array_Size]*bucket
}

// bucket is a linked list in each slot of the array
type bucket struct {
	head *bucketNode
}

// bucketNode is a node in a linked list that holds the value
type bucketNode struct {
	key  string
	next *bucketNode
}

//Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)

}

//Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	result := h.array[index].search(key)
	return result

}

//Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

//insert
func (b *bucket) insert(str string) {
	//key does not exist
	if !b.search(str) {
		newNode := &bucketNode{key: str}

		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Element already exists")
	}

}

//search will take in a key and return true if bucket has that key
func (b *bucket) search(str string) bool {
	currNode := b.head

	for currNode != nil {
		if currNode.key == str {
			return true
		}
		currNode = currNode.next
	}

	return false
}

//delete will take in a key and remove the node from the bucket
func (b *bucket) delete(str string) {

	//edge case for head node
	if b.head.key == str {
		b.head = b.head.next
		return
	}

	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == str {
			//delete
			prevNode.next = prevNode.next.next
		}
	}
}

//hash: get the ascii code for each char in key, sum it, divide it by array size, get remainder
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v) //int()gets ascii value of char
	}
	return sum % Array_Size
}

//Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}

	//make a bucket for each slot
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	myHashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"Kyle",
		"Stan",
		"Butters",
		"Randy",
		"Tolkien",
	}
	for _, v := range list {
		myHashTable.Insert(v)
	}
	myHashTable.Delete("Randy")

	for _, v := range myHashTable.array {
		fmt.Println(v)
	}
}
