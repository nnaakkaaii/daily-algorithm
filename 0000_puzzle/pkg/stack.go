package pkg

type Node struct {
	ID    int
	Field Field
	Key   int
}

type Stack struct {
	Nodes []*Node
}

var _ List = (*Stack)(nil)

func NewStack() *Stack {
	return &Stack{
		Nodes: make([]*Node, 0),
	}
}

func (s *Stack) Push(x interface{}) {
	n, ok := x.(*Node)
	if !ok {
		panic("unknown type")
	}
	s.Nodes = append(s.Nodes, n)
}

func (s *Stack) Pop() interface{} {
	if len(s.Nodes) == 0 {
		return nil
	}
	ind := len(s.Nodes) - 1
	n := s.Nodes[ind]
	s.Nodes = s.Nodes[:ind]
	return n
}

func (s *Stack) IsEmpty() bool {
	return len(s.Nodes) == 0
}

func (s *Stack) Len() int {
	return len(s.Nodes)
}

func (s *Stack) Less(i, j int) bool {
	return s.Nodes[i].Key < s.Nodes[j].Key
}

func (s *Stack) Swap(i, j int) {
	s.Nodes[i], s.Nodes[j] = s.Nodes[j], s.Nodes[i]
}
