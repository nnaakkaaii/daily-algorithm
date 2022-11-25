package pkg

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSimpleSearch_Search(t *testing.T) {
	t.Parallel()

	type args struct {
		list  List
		field Field
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"successful case: ",
			args{
				NewStack(),
				*NewField(
					[][]int{
						{1, 2, 3, 4},
						{6, 7, 8, 0},
						{5, 10, 11, 12},
						{9, 13, 14, 15},
					},
					4,
				),
				20,
			},
			8,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			d := NewAStar()

			if diff := cmp.Diff(tc.want, d.IterativeSearch(tc.args.f)); diff != "" {
				t.Error(diff)
			}
		})
	}
}