package lc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	SizeErrStr    = "size err"
	ValErr        = "val err"
	HeadValErrStr = "head val err"
	TailValErrStr = "tail val err"
)

func TestConstructor(t *testing.T) {
	l := Constructor()
	fmt.Println(l.size)
	assert.Equal(t, l.size, 0, SizeErrStr)
}

func TestMyLinkedList_AddAtHead(t *testing.T) {
	l := Constructor()
	l.AddAtHead(2)
	assert.Equal(t, l.size, 1, SizeErrStr)
	assert.Equal(t, l.head.val, 2, HeadValErrStr)
}

func TestMyLinkedList_AddAtTail(t *testing.T) {
	l := Constructor()
	l.AddAtTail(2)
	assert.Equal(t, l.size, 1, SizeErrStr)
	assert.Equal(t, l.head.val, 2, HeadValErrStr)
}

func TestMyLinkedList_SliceToLinkedList(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	l := Constructor()
	l.SliceToLinkedList(list)
	assert.Equal(t, len(list), l.size, SizeErrStr)
	assert.Equal(t, list[0], l.head.val, HeadValErrStr)
	assert.Equal(t, list[len(list)-1], l.tail.val, TailValErrStr)
}

func TestMyLinkedList_AddAtIndex(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	l := Constructor()
	l.SliceToLinkedList(list)
	assert.Equal(t, len(list), l.size, SizeErrStr)
	l.AddAtIndex(3, 33)
	assert.Equal(t, l.head.next.next.next.val, 33, ValErr)
}

func TestMyLinkedList_DeleteAtIndex(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	l := Constructor()
	l.SliceToLinkedList(list)
	assert.Equal(t, len(list), l.size, SizeErrStr)
	l.DeleteAtIndex(3)
	assert.Equal(t, l.head.next.next.next.val, 5, ValErr)
}

func TestMyLinkedList_DeleteAtHead(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	l := Constructor()
	l.SliceToLinkedList(list)
	assert.Equal(t, len(list), l.size, SizeErrStr)
	l.DeleteAtHead()
	assert.Equal(t, l.head.val, 2, HeadValErrStr)
}

func TestMyLinkedList_DeleteAtTail(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	l := Constructor()
	l.SliceToLinkedList(list)
	assert.Equal(t, len(list), l.size, SizeErrStr)
	l.DeleteAtTail()
	assert.Equal(t, l.tail.val, 5, HeadValErrStr)
}
