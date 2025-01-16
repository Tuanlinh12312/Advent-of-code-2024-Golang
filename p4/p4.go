package main

import (
	// "container/heap"
	// "container/list"

	// "slices"
	// "sort"

	"bufio"
	"fmt"
	"os"

	// "regexp"

	// "io"

	"strconv"
	"strings"
	// "regexp"
	// "math"
	// "cmp"
)

var (
	w *bufio.Writer
	r *bufio.Reader
)

func fScan(a ...any) (int, error) {
	return fmt.Fscan(r, a...)
}

func fScanf(format string, a ...any) (int, error) {
	return fmt.Fscanf(r, format, a...)
}

func fScanln(a ...any) (int, error) {
	return fmt.Fscanln(r, a...)
}

func fReadInts() ([]int, error) {
	inp, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	ans := make([]int, 0)
	for _, num := range strings.Fields(inp) {
		n, _ := strconv.Atoi(num)
		ans = append(ans, n)
	}

	return ans, nil
}

func fPrintf(format string, a ...any) (int, error) {
	return fmt.Fprintf(w, format, a...)
}

func fPrintln(a ...any) (int, error) {
	return fmt.Fprintln(w, a...)
}

func main() {
	// r = bufio.NewReader(os.Stdin)
	// w = bufio.NewWriter(os.Stdout)

	fin, _ := os.Open("input.txt")
	defer fin.Close()
	r = bufio.NewReader(fin)

	fout, _ := os.Create("output.txt")
	defer fout.Close()
	w = bufio.NewWriter(fout)

	defer w.Flush()

	a := make([]string, 0)

	for {
		var s string
		_, err := fScanln(&s)
		if err != nil {
			break
		}

		a = append(a, s)
	}

	ans := 0
	n, m := len(a), len(a[0])
	directions := [][2]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	count := func(i, j int) int {
		if a[i][j] != 'A' {
			return 0
		}

		S, M := 0, 0
		for k := 0; k < 4; k++ {
			in, jn := i+directions[k][0], j+directions[k][1]
			if a[in][jn] == 'S' {
				S++
			}
			if a[in][jn] == 'M' {
				M++
			}
		}

		if S != 2 || M != 2 {
			return 0
		}
		if a[i-1][j-1] == a[i+1][j+1] {
			return 0
		}
		return 1
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			ans += count(i, j)
		}
	}

	fPrintln(ans)
}
