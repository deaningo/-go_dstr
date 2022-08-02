package go_dstr

import (
	//"fmt"
	"sync"
)

type SingleObject interface{}

type SingleNode struct {
	Data SingleObject
	Next *SingleObject
}

//定义一个单链表,头部，尾部
type SingleList struct {
	mutex *sync.RWMutex
	Head  *SingleNode
	Tail  *SingleNode
	Size  uint
}

//初始化单链表
func (list *SingleList) Init() {
	list.Head = nil
	list.Tail = nil
	list.Size = 0
	list.mutex = new(sync.RWMutex)
}

//添加节点到链表的末尾
func (list *SingleList) Append(node *SingleNode) bool {
	if node == nil {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()

	if list.Size == 0 {
		list.Head = node
		list.Tail = node
		list.Size = 1
		return true
	}

	tail := list.Tail
	tail.Next = node
	list.Tail = node
	list.Size += 1
	return true
}
