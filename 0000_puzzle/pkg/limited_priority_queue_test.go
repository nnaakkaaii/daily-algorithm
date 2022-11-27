package pkg

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLimitedPriorityQueue_Scenario(t *testing.T) {
	t.Parallel()

	t.Run("scenario test", func(t *testing.T) {
		t.Parallel()

		pq := NewLimitedPriorityQueue(3)

		pq.Push(&Node{1, Field{}, 3})
		pq.Push(&Node{2, Field{}, 4})
		pq.Push(&Node{3, Field{}, 2})
		pq.Push(&Node{4, Field{}, 1})
		pq.Push(&Node{5, Field{}, 5})

		v := make([]int, 0)
		for {
			if pq.IsEmpty() {
				break
			}
			n, ok := pq.Pop().(*Node)
			if !ok {
				t.Fatal("unknown type")
			}
			v = append(v, n.ID)
		}

		if diff := cmp.Diff([]int{1, 2, 5}, v); diff != "" {
			t.Error(diff)
		}
	})
}
