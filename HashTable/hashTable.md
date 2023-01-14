hash tables use a hash function to generate the index for an element
collision handeling :
(open handeling) go to the next index, to search,start traeing from the orignal index
benefits of open handling deminish as collsions grow
(seperate chaining): store multiple element in one index by use of a linked list
each index in hashset hold a pointer to the head of a LL and the LL can grow dynamicaly

(seperate chaining):
keys are the values in the LL
buckets are the LL in each index of the array

insert():
O(1) in average case
O(N) in worst case id only collisions
delete():
O(1) in average case
search():
O(1) in average case
 
Hash Table Part (array):
structure - HashTable
method - Insert()
method - Search()
method - Delete()

Bucket Part (Linked List):
structure - bucket
structure - bucketNode
method - Insert()
method - Search()
method - Delete()



