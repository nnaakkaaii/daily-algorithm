package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/nnaakkaaii/daily-algorithm/0000_puzzle/pkg"
)

const (
	FieldSize = 3
	MaxDepth  = 100
)

type Search string

const (
	SimpleSearch       Search = "simple-search"
	IterativeDeepening Search = "iterative-deepening"
)

var searches = map[Search]func(pkg.List) pkg.Search{
	SimpleSearch:       pkg.NewSimpleSearch,
	IterativeDeepening: pkg.NewIterativeDeepening,
}

type List string

const (
	Stack         List = "stack"
	Queue         List = "queue"
	PriorityQueue List = "priority-queue"
)

var lists = map[List]pkg.List{
	Stack:         pkg.NewStack(),
	Queue:         pkg.NewQueue(),
	PriorityQueue: pkg.NewPriorityQueue(),
}

func main() {
	var search = flag.String(
		"search",
		"simple-search",
		"simple-search|iterative-search",
	)
	var list = flag.String(
		"list",
		"priority-queue",
		"stack|queue|priority-queue",
	)
	flag.Parse()

	realMain(os.Stdout, os.Stdin, Search(*search), List(*list))
}

func realMain(w io.Writer, r io.Reader, search Search, list List) {
	br := bufio.NewReader(r)
	bw := bufio.NewWriter(w)
	defer bw.Flush()

	mat := make([][]int, 0, FieldSize)
	for i := 0; i < FieldSize; i++ {
		row := make([]int, 0, FieldSize)
		for j := 0; j < FieldSize; j++ {
			var e int
			fmt.Fscan(br, &e)
			row = append(row, e)
		}
		mat = append(mat, row)
	}

	f := pkg.NewField(mat, FieldSize)
	result := searches[search](lists[list]).Search(*f, MaxDepth)

	fmt.Fprintln(bw, result)
}
