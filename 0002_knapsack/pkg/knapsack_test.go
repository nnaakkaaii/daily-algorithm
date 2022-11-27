package pkg

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestKnapsack_Solve(t *testing.T) {
	t.Parallel()

	type args struct {
		values    []int
		weights   []int
		maxWeight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "successful case",
			args: args{
				values:    []int{4, 5, 2, 8},
				weights:   []int{2, 2, 1, 3},
				maxWeight: 5,
			},
			want: 13,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			x := NewKnapsack(tc.args.values, tc.args.weights)
			actual := x.Solve(tc.args.maxWeight)

			if diff := cmp.Diff(tc.want, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}
