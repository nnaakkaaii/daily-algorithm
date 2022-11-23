package pkg

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIDAStar_Search(t *testing.T) {
	t.Parallel()

	d := NewIDAStar()

	f := NewField([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 0, 12},
		{13, 14, 11, 15},
	}, 4)

	r1 := d.Search(*f, 1)
	if diff := cmp.Diff(-1, r1); diff != "" {
		t.Error(diff)
	}

	r2 := d.Search(*f, 2)
	if diff := cmp.Diff(2, r2); diff != "" {
		t.Error(diff)
	}
}

func TestIDAStar_IterativeSearch(t *testing.T) {
	t.Parallel()

	d := NewIDAStar()

	f := NewField([][]int{
		{1, 2, 3, 4},
		{6, 7, 8, 0},
		{5, 10, 11, 12},
		{9, 13, 14, 15},
	}, 4)

	r1 := d.IterativeSearch(*f)
	if diff := cmp.Diff(8, r1); diff != "" {
		t.Error(diff)
	}
}
