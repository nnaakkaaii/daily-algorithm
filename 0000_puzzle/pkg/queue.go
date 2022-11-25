package pkg

type Queue struct {
	Nodes []*Node
}

var _ List = (*Queue)(nil)

func NewQueue() *Queue {
	return &Queue{
		Nodes: make([]*Node, 0),
	}
}

func (q *Queue) Push(x interface{}) {
	n, ok := x.(*Node)
	if !ok {
		panic("unknown type")
	}
	q.Nodes = append(q.Nodes, n)
}

func (q *Queue) Pop() interface{} {
	if len(q.Nodes) == 0 {
		return nil
	}
	n := q.Nodes[0]
	q.Nodes = q.Nodes[1:]
	return n
}

func (q *Queue) IsEmpty() bool {
	return len(q.Nodes) == 0
}
