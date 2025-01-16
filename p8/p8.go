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
		_, err := fScanln(&s)
		if err != nil {
			break
		}
		a = append(a, []rune(s))
	}

	ans := 0
	n, m := len(a), len(a[0])
	check := make(map[[2]int]bool)
	pos := make(map[rune][][2]int)

	for i := 0; i < n; i ++ {
		for j := 0; j < m; j ++ {
			if a[i][j] != '.' {
				pos[a[i][j]] = append(pos[a[i][j]], [2]int{i, j})
			}
		}
	} 

	AddNode := func(x, y int) bool {
		if x < 0 || x >= n || y < 0 || y >= m {
			return true
		}
		if !check[[2]int{x, y}] {
			ans ++
		}
		if a[x][y] == '.' {
			a[x][y] = '#'
		}
		check[[2]int{x, y}] = true
		return false
	}

	for _, a := range pos {
		k := len(a)
		for i := 0; i < k; i ++ {
			for j := i + 1; j < k; j ++ {
				for m := 0;; m++ {
					if AddNode(a[i][0] * (m + 1) - a[j][0] * m, a[i][1] * (m + 1) - a[j][1]*m) {
						break
					}
				}
				for m := 0;; m++ {
					if AddNode(a[j][0] * (m + 1) - a[i][0] * m, a[j][1] * (m + 1) - a[i][1]*m) {
						break
					}
				}
			}
		}
	}

	fPrintln(ans)
}
