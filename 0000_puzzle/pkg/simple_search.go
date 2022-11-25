package pkg

type SimpleSearch struct {
	Memory List
}

func NewSimpleSearch(l List) Search {
	return &SimpleSearch{
		Memory: l,
	}
}

func (s *SimpleSearch) Search(f Field, limit int) int {
	k := f.S + f.Manhattan() - 1
	for _, id := range f.Swappable() {
		s.Memory.Push(&Node{id, f.Copy(), k})
	}

	for {
		if f.Manhattan() == 0 {
			return f.S
		}

		if s.Memory.IsEmpty() {
			return -1
		}

		n, ok := s.Memory.Pop().(*Node)
		if !ok {
			panic("unknown type")
		}

		f = n.Field
		f.Swap(n.ID)

		k = f.S + f.Manhattan() - 1
		if k >= limit {
			continue
		}
		for _, id := range f.Swappable() {
			if id == n.ID {
				continue
			}
			s.Memory.Push(&Node{id, f.Copy(), k})
		}
	}
}
