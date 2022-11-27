package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/nnaakkaaii/daily-algorithm/0002_knapsack/pkg"
)

func main() {
	realMain(os.Stdout, os.Stdin)
}

func realMain(w io.Writer, r io.Reader) {
	br := bufio.NewReader(r)
	bw := bufio.NewWriter(w)
	defer bw.Flush()

	var n int
	fmt.Fscan(br, &n)

	var mw int
	fmt.Fscan(br, &mw)

	vs := make([]int, 0, n)
	ws := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(br, &x)
		vs = append(vs, x)
		fmt.Fscan(br, &x)
		ws = append(ws, x)
	}

	result := pkg.NewKnapsack(vs, ws).Solve(mw)

	fmt.Fprintln(bw, result)
}
