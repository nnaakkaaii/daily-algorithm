package pkg

type IDAStar struct {
	Memory *Stack
}

func NewIDAStar() SearchAlgorithm {
	return &IDAStar{Memory: NewStack()}
}

func (d *IDAStar) Search(f Field, limit int) int {
	for _, id := range f.Swappable() {
		d.Memory.Push(id, f.Copy())
	}

	for {
		if f.Manhattan() == 0 {
			return f.S
		}
		n := d.Memory.Pop()
		if n == nil {
			return -1
		}
		f = n.Field
		f.Swap(n.ID)
		if f.S+f.Manhattan()-1 >= limit {
			continue
		}
		for _, id := range f.Swappable() {
			if id == n.ID {
				continue
			}
			d.Memory.Push(id, f.Copy())
		}
	}
}

func (d *IDAStar) IterativeSearch(f Field) int {
	i := 1
	for {
		r := d.Search(f, i)
		if r != -1 {
			return r
		}
		i *= 2
	}
}
