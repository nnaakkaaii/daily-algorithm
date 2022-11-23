package pkg

type AStar struct {
	Memory *PriorityQueue
}

func NewAStar() SearchAlgorithm {
	return &AStar{Memory: NewPriorityQueue()}
}

func (d *AStar) IterativeSearch(f Field) int {
	k := f.S + f.Manhattan() - 1
	for _, id := range f.Swappable() {
		d.Memory.Push(id, f.Copy(), k)
	}

	for {
		if f.Manhattan() == 0 {
			return f.S
		}
		if d.Memory.IsEmpty() {
			return -1
		}
		n := d.Memory.Pop()
		f = n.Field
		f.Swap(n.ID)
		k = f.S + f.Manhattan() - 1
		for _, id := range f.Swappable() {
			if id == n.ID {
				continue
			}
			d.Memory.Push(id, f.Copy(), k)
		}
	}
}
