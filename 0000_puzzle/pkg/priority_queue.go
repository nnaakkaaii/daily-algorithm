package pkg

import (
	"container/heap"
)

type NodeP struct {
	ID    int
	Field Field
	Key   int
}

type StackP struct {
	Nodes []*NodeP
}

func NewStackP() *StackP {
	return &StackP{Nodes: make([]*NodeP, 0)}
}

func (s *StackP) Len() int {
	return len(s.Nodes)
}

func (s *StackP) Less(i, j int) bool {
	return s.Nodes[i].Key < s.Nodes[j].Key
}

func (s *StackP) Swap(i, j int) {
	s.Nodes[i], s.Nodes[j] = s.Nodes[j], s.Nodes[i]
}

func (s *StackP) Push(x interface{}) {
	n, ok := x.(*NodeP)
	if !ok {
		panic("unknown type")
	}
	s.Nodes = append(s.Nodes, n)
}

func (s *StackP) Pop() interface{} {
	ind := len(s.Nodes) - 1
	n := s.Nodes[ind]
	s.Nodes = s.Nodes[:ind]
	return n
}

func (s *StackP) IsEmpty() bool {
	return len(s.Nodes) == 0
}

type PriorityQueue struct {
	Stack *StackP
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{Stack: NewStackP()}
}

func (pq *PriorityQueue) Push(id int, f Field, k int) {
	heap.Push(pq.Stack, &NodeP{id, f, k})
}

func (pq *PriorityQueue) Pop() *NodeP {
	x := heap.Pop(pq.Stack)
	n, ok := x.(*NodeP)
	if !ok {
		return nil
	}
	return n
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Stack.IsEmpty()
}
