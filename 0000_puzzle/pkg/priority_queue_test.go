package pkg

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPriorityQueue_Scenario(t *testing.T) {
	t.Parallel()

	t.Run("scenario test", func(t *testing.T) {
		t.Parallel()

		pq := NewPriorityQueue()

		pq.Push(1, Field{}, 3)
		pq.Push(2, Field{}, 4)
		pq.Push(3, Field{}, 2)
		pq.Push(4, Field{}, 1)
		pq.Push(5, Field{}, 5)

		v := make([]int, 0)
		for {
			if pq.IsEmpty() {
				break
			}
			v = append(v, pq.Pop().ID)
		}

		if diff := cmp.Diff([]int{4, 3, 1, 2, 5}, v); diff != "" {
			t.Error(diff)
		}
	})
}
