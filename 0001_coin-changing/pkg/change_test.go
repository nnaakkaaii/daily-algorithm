package pkg

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestChange_Solve(t *testing.T) {
	t.Parallel()

	type args struct {
		coins []int
		total int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "successful case",
			args: args{
				coins: []int{1, 2, 7, 8, 12, 50},
				total: 15,
			},
			want: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			x := NewChange(tc.args.coins)
			actual := x.Solve(tc.args.total)

			if diff := cmp.Diff(tc.want, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}
