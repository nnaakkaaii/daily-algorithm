package pkg

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestQueue_Scenario(t *testing.T) {
	t.Parallel()

	t.Run("scenario test", func(t *testing.T) {
		t.Parallel()

		q := NewQueue()

		q.Push(&Node{1, Field{}, 3})
		q.Push(&Node{2, Field{}, 4})
		q.Push(&Node{3, Field{}, 2})
		q.Push(&Node{4, Field{}, 1})
		q.Push(&Node{5, Field{}, 5})

		v := make([]int, 0)
		for {
			emp := q.IsEmpty()
			x := q.Pop()

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

		if diff := cmp.Diff([]int{1, 2, 3, 4, 5}, v); diff != "" {
			t.Error(diff)
		}
	})
}
