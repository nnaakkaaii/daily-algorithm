package pkg

import (
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Field struct {
	X  map[int]Coordinate
	IX map[Coordinate]int
	T  map[int]Coordinate
	D  int
	S  int
}

func InitialCoordinates(d int) map[int]Coordinate {
	t := map[int]Coordinate{}
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			t[(1+i*d+j)%(d*d)] = Coordinate{j, i}
		}
	}
	return t
}

func Array2Coordinates(a [][]int) map[int]Coordinate {
	t := map[int]Coordinate{}
	for i, row := range a {
		for j, e := range row {
			t[e] = Coordinate{j, i}
		}
	}
	return t
}

func InverseCoordinates(c map[int]Coordinate) map[Coordinate]int {
	t := map[Coordinate]int{}
	for k, v := range c {
		t[v] = k
	}
	return t
}

func NewField(x [][]int, d int) *Field {
	xc := Array2Coordinates(x)
	return &Field{
		xc,
		InverseCoordinates(xc),
		InitialCoordinates(d),
		d,
		0,
	}
}

func (f *Field) Manhattan() int {
	d := 0
	for k := range f.T {
		d += int(math.Abs(float64(f.X[k].X - f.T[k].X)))
		d += int(math.Abs(float64(f.X[k].Y - f.T[k].Y)))
	}
	return d / 2
}

func (f *Field) Swappable() []int {
	c := f.X[0]
	re := make([]int, 0, 4)
	if c.X-1 >= 0 {
		r := f.IX[Coordinate{c.X - 1, c.Y}]
		re = append(re, r)
	}
	if c.Y-1 >= 0 {
		r := f.IX[Coordinate{c.X, c.Y - 1}]
		re = append(re, r)
	}
	if c.X+1 < f.D {
		r := f.IX[Coordinate{c.X + 1, c.Y}]
		re = append(re, r)
	}
	if c.Y+1 < f.D {
		r := f.IX[Coordinate{c.X, c.Y + 1}]
		re = append(re, r)
	}
	return re
}

func (f *Field) Swap(i int) {
	ci, c := f.X[i], f.X[0]
	f.X[0], f.X[i] = ci, c
	f.IX[c], f.IX[ci] = f.IX[ci], f.IX[c]
	f.S += 1
}

func (f *Field) Render() string {
	r := ""
	ml := len(strconv.Itoa(f.D * f.D))
	for i := 0; i < f.D; i++ {
		for j := 0; j < f.D; j++ {
			s := strconv.Itoa(f.IX[Coordinate{j, i}])
			r += s + strings.Repeat(" ", 1+ml-len(s))
		}
		r += "\n"
	}
	return r
}

func (f *Field) Copy() Field {
	X := map[int]Coordinate{}
	IX := map[Coordinate]int{}
	T := map[int]Coordinate{}

	for k, v := range f.X {
		X[k] = v
	}
	for k, v := range f.IX {
		IX[k] = v
	}
	for k, v := range f.T {
		T[k] = v
	}

	return Field{
		X:  X,
		IX: IX,
		T:  T,
		D:  f.D,
		S:  f.S,
	}
}
