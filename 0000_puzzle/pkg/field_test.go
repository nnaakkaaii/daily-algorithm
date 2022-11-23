package pkg

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArray2Coordinates(t *testing.T) {
	t.Parallel()

	type args struct {
		a [][]int
	}
	tests := []struct {
		name string
		args args
		want map[int]Coordinate
	}{
		{
			name: "successcul case",
			args: args{
				[][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
					{13, 14, 15, 0},
				},
			},
			want: map[int]Coordinate{
				1:  {0, 0},
				2:  {1, 0},
				3:  {2, 0},
				4:  {3, 0},
				5:  {0, 1},
				6:  {1, 1},
				7:  {2, 1},
				8:  {3, 1},
				9:  {0, 2},
				10: {1, 2},
				11: {2, 2},
				12: {3, 2},
				13: {0, 3},
				14: {1, 3},
				15: {2, 3},
				0:  {3, 3},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			act := Array2Coordinates(tc.args.a)

			if diff := cmp.Diff(tc.want, act); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestInitialCoordinates(t *testing.T) {
	t.Parallel()

	type args struct {
		d int
	}
	tests := []struct {
		name string
		args args
		want map[int]Coordinate
	}{
		{
			"successful case",
			args{4},
			map[int]Coordinate{
				1:  {0, 0},
				2:  {1, 0},
				3:  {2, 0},
				4:  {3, 0},
				5:  {0, 1},
				6:  {1, 1},
				7:  {2, 1},
				8:  {3, 1},
				9:  {0, 2},
				10: {1, 2},
				11: {2, 2},
				12: {3, 2},
				13: {0, 3},
				14: {1, 3},
				15: {2, 3},
				0:  {3, 3},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			act := InitialCoordinates(tc.args.d)

			if diff := cmp.Diff(tc.want, act); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestField_Render(t *testing.T) {
	t.Parallel()

	f := NewField(
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 0, 12},
			{13, 14, 11, 15},
		},
		4,
	)
	if diff := cmp.Diff("1  2  3  4  \n5  6  7  8  \n9  10 0  12 \n13 14 11 15 \n", f.Render()); diff != "" {
		t.Error(diff)
	}
}

func TestField_Scenario(t *testing.T) {
	t.Parallel()

	f := NewField(
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 0, 12},
			{13, 14, 11, 15},
		},
		4,
	)
	if diff := cmp.Diff([]int{10, 7, 12, 11}, f.Swappable()); diff != "" {
		t.Error(diff)
	}
	if diff := cmp.Diff(2, f.Manhattan()); diff != "" {
		t.Error(diff)
	}
	f.Swap(11)
	if diff := cmp.Diff([]int{14, 11, 15}, f.Swappable()); diff != "" {
		t.Error(diff)
	}
	if diff := cmp.Diff(1, f.Manhattan()); diff != "" {
		t.Error(diff)
	}
	f.Swap(15)
	if diff := cmp.Diff([]int{15, 12}, f.Swappable()); diff != "" {
		t.Error(diff)
	}
	if diff := cmp.Diff(0, f.Manhattan()); diff != "" {
		t.Error(diff)
	}
}
