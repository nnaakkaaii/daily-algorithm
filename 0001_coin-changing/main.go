package main

import (
	"bufio"
	"fmt"
	"github.com/nnaakkaaii/daily-algorithm/0001_coin-changing/pkg"
	"io"
	"os"
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

	var m int
	fmt.Fscan(br, &m)

	coins := make([]int, 0, m)
	for i := 0; i < m; i++ {
		var x int
		fmt.Fscan(br, &x)
		coins = append(coins, x)
	}

	result := pkg.NewChange(coins).Solve(n)

	fmt.Fprintln(bw, result)
}
