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

func fGetln(s *string) error {
	sn, err := r.ReadString('\n')
	if err != nil {
		return err
	}
	*s = sn 
	return nil
}

func SplitInts(s string, sep rune) ([]int, error) {
	s = strings.TrimSpace(s)

	ans := make([]int, 0)
	for _, num := range strings.FieldsFunc(s, func(r rune) bool {
		return r == sep
	}) {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}

		ans = append(ans, n)
	}

	return ans, nil
}

func fReadInts(sep rune) ([]int, error) {
	inp, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	
	return SplitInts(inp, sep)
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
		_, err := fScan(&s) 
		if err != nil {
			break
		}
		a = append(a, []rune(s))
	}

	ans := 0
	n, m := len(a), len(a[0])
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := 0; i < n; i ++ {
		for j := 0; j < m; j ++ {
			if a[i][j] != '0' {
				continue
			}

			pos := make(map[[2]int]int, 0)
			pos[[2]int{i, j}] = 1

			for c := '1'; c <= '9'; c ++ {
				npos := make(map[[2]int]int, 0)
				for key, val := range pos {
					x, y := key[0], key[1]
					for k := 0; k < len(directions); k ++ {
						xn, yn := x + directions[k][0], y + directions[k][1]
						if xn < 0 || xn >= n || yn < 0 || yn >= m {
							continue
						}
						if a[xn][yn] == c {
							npos[[2]int{xn, yn}] += val
						}
					}
				}

				pos = npos
			}

			for _, val := range pos	{
				ans += val
			}
		}
	}

	fPrintln(ans)
}
