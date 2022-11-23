package pkg

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIterativeDeepening_Search(t *testing.T) {
	t.Parallel()

	type args struct {
		f Field
		l int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"reach limit",
			args{
				*NewField(
					[][]int{
						{1, 2, 3, 4},
						{5, 6, 7, 8},
						{9, 10, 0, 12},
						{13, 14, 11, 15},
					},
					4,
				),
				1,
			},
			-1,
		},
		{
			"success",
			args{
				*NewField(
					[][]int{
						{1, 2, 3, 4},
						{5, 6, 7, 8},
						{9, 10, 0, 12},
						{13, 14, 11, 15},
					},
					4,
				),
				2,
			},
			2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			d := NewIterativeDeepening()

			if diff := cmp.Diff(tc.want, d.Search(tc.args.f, tc.args.l)); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestIterativeDeepening_IterativeSearch(t *testing.T) {
	t.Parallel()

	type args struct {
		f Field
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"successful case",
			args{
				*NewField(
					[][]int{
						{1, 2, 3, 4},
						{6, 7, 8, 0},
						{5, 10, 11, 12},
						{9, 13, 14, 15},
					},
					4,
				),
			},
			8,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			d := NewIDAStar()

			if diff := cmp.Diff(tc.want, d.IterativeSearch(tc.args.f)); diff != "" {
				t.Error(diff)
			}
		})
	}
}
