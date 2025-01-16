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

func fReadInts(sep rune) ([]int, error) {
	inp, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	inp = strings.TrimSpace(inp)

	ans := make([]int, 0)
	for _, num := range strings.FieldsFunc(inp, func(r rune) bool {
		return r == sep
	}) {
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
	
	a := make([][]rune, 0)
	for {
		var s string 
		_, err := fScanln(&s)
		if err != nil {
			break
		}
		a = append(a, []rune(s))
	}

	n, m := len(a), len(a[0])
	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	
	sti, stj := 0, 0
	for i := 0; i < n; i ++ {
		for j := 0; j < m; j ++ {
			if a[i][j] == '^' {
				sti, stj = i, j
			}
		}
	}
	
	ans := 0
	for i := 0; i < n; i ++ {
		for j := 0; j < m; j ++ {
			if a[i][j] == '.' {
				a[i][j] = '#'
				check := make(map[[3]int]bool)

				var dfs func(int, int, int) bool
				dfs = func(x, y, d int) bool {
					if check[[3]int{x, y, d}] {
						return true
					}
					xn, yn, dn := x + directions[d][0], y + directions[d][1], (d + 1)%4
					check[[3]int{x, y, d}] = true
			
					if xn < 0 || xn >= n || yn < 0 || yn >= m {
						return false
					} else if a[xn][yn] == '#' {
						return dfs(x, y, dn)
					} else {
						return dfs(xn, yn, d)
					}
				}

				if dfs(sti, stj, 0) {
					ans ++
				}
				a[i][j] = '.'
			}
		}
	}
	fPrintln(ans)
}
