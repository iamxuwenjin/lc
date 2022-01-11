package lc

import (
	"sync"
)

// MyLinkedList
// @Description: 链表
type MyLinkedList struct {
	head *node
	tail *node
	mu   sync.Mutex
	size int
}

type node struct {
	val  int
	pre  *node
	next *node
}

// Constructor
// return MyLinkedList
func Constructor() MyLinkedList {
	return MyLinkedList{size: 0}
}

// Get
// param index offset of linked list
// return int val
func (l *MyLinkedList) Get(index int) int {
	if index == 0 {
		return l.head.val
	}
	if index >= l.size || l.size == 0 {
		// TODO err
		return -1
	}

	var val int
	cur := l.head
	for i := 1; i <= index; i++ {
		cur = cur.next
		if i == index {
			val = cur.val
		}
	}
	return val
}

// AddAtHead 从头部增加节点
// param val
func (l *MyLinkedList) AddAtHead(val int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.addAtHead(val)
}

// AddAtHead 从头部增加节点 private
// param val
func (l *MyLinkedList) addAtHead(val int) {
	n := &node{
		val:  val,
		pre:  nil,
		next: l.head,
	}
	if l.head != nil {
		l.head.pre = n
	}
	l.head = n
	if l.tail == nil {
		l.tail = n
	}
	l.size++
}

// AddAtTail 从尾部增加节点
// param val
func (l *MyLinkedList) AddAtTail(val int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.addAtTail(val)
}

// addAtTail 从尾部增加节点 private
// param val
func (l *MyLinkedList) addAtTail(val int) {
	n := &node{
		val:  val,
		pre:  l.tail,
		next: nil,
	}
	if l.tail != nil {
		l.tail.next = n
	}
	l.tail = n
	if l.head == nil {
		l.head = n
	}
	l.size++
}

// DeleteAtHead 删除头部节点
func (l *MyLinkedList) DeleteAtHead() {
	if l.size == 0 {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.deleteAtHead()

}

// DeleteAtHead 删除头部节点 private
func (l *MyLinkedList) deleteAtHead() {
	next := l.head.next
	// 链表只有一个节点
	if next == nil {
		l.head = nil
		l.tail = nil
		l.size--
		return
	}
	next.pre = nil
	l.head = next
	l.size--
}

// DeleteAtTail 删除尾部节点
func (l *MyLinkedList) DeleteAtTail() {
	if l.size == 0 {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.deleteAtTail()
}

// DeleteAtTail 删除尾部节点 private
func (l *MyLinkedList) deleteAtTail() {
	pre := l.tail.pre
	if pre == nil {
		l.head = nil
		l.tail = nil
		l.size--
		return
	}
	pre.next = nil
	l.tail = pre
	l.size--
}

// AddAtIndex 指定位置增加节点
// param index
// param val
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if index == 0 {
		l.addAtHead(val)
		return
	}
	if index >= l.size {
		l.addAtTail(val)
		return
	}
	cur := l.head
	var pre *node
	for i := 1; i <= index; i++ {
		pre = cur
		cur = cur.next
		if i == index {
			n := &node{
				val:  val,
				pre:  cur.pre,
				next: cur,
			}
			pre.next = n
		}
	}
	l.size++
}

// DeleteAtIndex 指定位置删除节点
// param index
func (l *MyLinkedList) DeleteAtIndex(index int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if index >= l.size {
		return
	}
	if index == 0 {
		l.deleteAtHead()
		return
	}
	if index == l.size-1 {
		l.deleteAtTail()
		return
	}
	var pre *node
	cur := l.head
	for i := 1; i <= index; i++ {
		pre = cur
		cur = cur.next
		if i == index {
			pre.next = cur.next
		}
	}
	l.size--
}

func (l *MyLinkedList) SliceToLinkedList(list []int) {
	for i := range list {
		l.AddAtIndex(i, list[i])
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
