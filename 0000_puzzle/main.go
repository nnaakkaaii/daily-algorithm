package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/nnaakkaaii/daily-algorithm/0000_puzzle/pkg"
	"io"
	"os"
)

const FieldSize = 4

type Algorithm string

const (
	IterativeDeepeningAlgorithm Algorithm = "IterativeSearch"
	IDAStarAlgorithm            Algorithm = "IDAStar"
)

var algorithms = map[Algorithm]pkg.SearchAlgorithm{
	IterativeDeepeningAlgorithm: pkg.NewIterativeDeepening(),
	IDAStarAlgorithm:            pkg.NewIDAStar(),
}

func main() {
	var algorithm = flag.String(
		"algorithm",
		"IDAStar",
		"IterativeDeepening|IDAStar",
	)
	flag.Parse()

	realMain(os.Stdout, os.Stdin, Algorithm(*algorithm))
}

func realMain(w io.Writer, r io.Reader, algorithm Algorithm) {
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
	result := algorithms[algorithm].IterativeSearch(*f)

	fmt.Fprintln(bw, result)
}
