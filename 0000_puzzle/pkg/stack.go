package pkg

type Node struct {
	ID    int
	Field Field
}

type Stack struct {
	Nodes []*Node
}

func NewStack() *Stack {
	return &Stack{[]*Node{}}
}

func (s *Stack) Push(id int, field Field) {
	s.Nodes = append(s.Nodes, &Node{id, field})
}

func (s *Stack) Pop() *Node {
	if len(s.Nodes) == 0 {
		return nil
	}
	ind := len(s.Nodes) - 1
	n := s.Nodes[ind]
	s.Nodes = s.Nodes[:ind]
	return n
}
