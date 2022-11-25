package pkg

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIterativeDeepening_Search(t *testing.T) {
	t.Parallel()

	type args struct {
		list  List
		field Field
		limit int
	}
	f := *NewField([][]int{{1, 2, 3}, {4, 0, 6}, {7, 5, 8}}, 3)

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"successful case: IDA Star",
			args{NewStack(), f, 5},
			2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewIterativeDeepening(tc.args.list)

			if diff := cmp.Diff(tc.want, s.Search(tc.args.field, tc.args.limit)); diff != "" {
				t.Error(diff)
			}
		})
	}
}
