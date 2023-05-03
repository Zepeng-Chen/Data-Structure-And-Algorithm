package ds

type SinglyListNode struct {
	val  int
	next *SinglyListNode
}

type DoublyLinkedNode struct {
	val  int
	prev *DoublyLinkedNode
	next *DoublyLinkedNode
}

type SinglyLinkedList struct {
	head *SinglyListNode
}

func Constructor() SinglyLinkedList {
	var head SinglyLinkedList
	return head
}

func (sl SinglyLinkedList) addAtHead(val int) {
	newNode := SinglyListNode{val: val}
	newNode.next = sl.head
	sl.head = &newNode
}

func (sl SinglyLinkedList) AddAtTail(val int) {
	for sl.head.next != nil {
		sl.head = sl.head.next
	}
	sl.head.next = &SinglyListNode{val: val}
}

func (sl SinglyLinkedList) AddAtIndex(index, val int) {
	for i := 0; i < index; i++ {
		if sl.head == nil {
			return
		}
		sl.head = sl.head.next
	}
	afterNode := sl.head.next
	newNode := SinglyListNode{val: val}
	sl.head.next = &newNode
	newNode.next = afterNode
}

func (sl SinglyLinkedList) DeleteAtIndex(index int) {
	for i := 0; i < index; i++ {
		if sl.head == nil {
			return
		}
		sl.head = sl.head.next
	}
	sl.head.next = sl.head.next.next
}
