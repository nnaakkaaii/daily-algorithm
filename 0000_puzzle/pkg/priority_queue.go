package pkg

import (
	"container/heap"
)

type PriorityQueue struct {
	Stack *Stack
}

var _ List = (*PriorityQueue)(nil)

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		Stack: NewStack(),
	}
}

func (pq *PriorityQueue) Push(x interface{}) {
	n, ok := x.(*Node)
	if !ok {
		panic("unknown type")
	}
	heap.Push(pq.Stack, n)
}

func (pq *PriorityQueue) Pop() interface{} {
	x := heap.Pop(pq.Stack)
	n, ok := x.(*Node)
	if !ok {
		return nil
	}
	return n
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Stack.IsEmpty()
}
