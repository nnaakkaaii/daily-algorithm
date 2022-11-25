package pkg

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestStack_Scenario(t *testing.T) {
	t.Parallel()

	t.Run("scenario test", func(t *testing.T) {
		t.Parallel()

		s := NewStack()

		s.Push(&Node{1, Field{}, 3})
		s.Push(&Node{2, Field{}, 4})
		s.Push(&Node{3, Field{}, 2})
		s.Push(&Node{4, Field{}, 1})
		s.Push(&Node{5, Field{}, 5})

		if diff := cmp.Diff(5, s.Len()); diff != "" {
			t.Error(diff)
		}
		if diff := cmp.Diff(false, s.Less(1, 2)); diff != "" {
			t.Error(diff)
		}

		v := make([]int, 0)
		for {
			if s.IsEmpty() {
				break
			}
			n, ok := s.Pop().(*Node)
			if !ok {
				t.Fatal("unknown type")
			}
			v = append(v, n.ID)
		}

		if diff := cmp.Diff([]int{5, 4, 3, 2, 1}, v); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("scenario test: swap", func(t *testing.T) {
		t.Parallel()

		s := NewStack()

		s.Push(&Node{1, Field{}, 3})
		s.Push(&Node{2, Field{}, 4})

		s.Swap(0, 1)

		v := make([]int, 0)
		for {
			emp := s.IsEmpty()
			x := s.Pop()

			if diff := cmp.Diff(x == nil, emp); diff != "" {
				t.Error(diff)
			}
			if x == nil {
				break
			}

			n, ok := x.(*Node)
			if !ok {
				t.Fatal("unknown type")
			}
			v = append(v, n.ID)
		}

		if diff := cmp.Diff([]int{1, 2}, v); diff != "" {
			t.Error(diff)
		}
	})
}
