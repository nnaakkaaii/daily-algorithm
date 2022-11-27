package pkg

import (
	"container/heap"
)

type LimitedPriorityQueue struct {
	BeamWidth int
	Stack     *Stack
}

var _ List = (*LimitedPriorityQueue)(nil)

func NewLimitedPriorityQueue(bw int) *LimitedPriorityQueue {
	return &LimitedPriorityQueue{
		BeamWidth: bw,
		Stack:     NewStack(),
	}
}

func (pq *LimitedPriorityQueue) Push(x interface{}) {
	n, ok := x.(*Node)
	if !ok {
		panic("unknown type")
	}

	if pq.IsFull() {
		nx, ok := pq.Pop().(*Node)
		if !ok {
			panic("unknown type")
		}
		// 既存のノードの方がkeyが大きい場合は戻す
		if nx.Key >= n.Key {
			n = nx
		}
	}

	heap.Push(pq.Stack, n)
}

func (pq *LimitedPriorityQueue) Pop() interface{} {
	x := heap.Pop(pq.Stack)
	n, ok := x.(*Node)
	if !ok {
		return nil
	}
	return n
}

func (pq *LimitedPriorityQueue) IsEmpty() bool {
	return pq.Stack.IsEmpty()
}

func (pq *LimitedPriorityQueue) IsFull() bool {
	return pq.Stack.Len() >= pq.BeamWidth
}
