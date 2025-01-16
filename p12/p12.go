package main

import (
	// "container/heap"
	"container/list"

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

	// read matrix
	mat := make([][]rune, 0)
	for {
		var s string
		_, err := fScan(&s)
		if err != nil {
			break
		}
		mat = append(mat, []rune(s))
	}

	ans := 0
	n, m := len(mat), len(mat[0])
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

	check := make([][]bool, n)
	for i := range check {
		check[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if check[i][j] {
				continue // skip visited nodes
			}

			q := list.New()
			area, perimeter := 0, 0 // perimeter = area * 4 - overlapping edges
			sides := 0              // sides = corners

			Corners := func(x, y int) int {
				// checks if the adjacent node in direction di is the same type
				IsSame := func(di int) bool {
					xn, yn := x+directions[di][0], y+directions[di][1]
					if xn < 0 || xn >= n || yn < 0 || yn >= m {
						return false
					}
					return mat[xn][yn] == mat[x][y]
				}

				ans := 0
				for di := 0; di < 4; di++ {
					if !IsSame(di) && !IsSame((di+1)%4) {
						ans++
					}
					if IsSame(di) && IsSame((di+1)%4) && !IsSame(di+4) {
						ans++
					}
				}
				return ans
			}

			Add := func(x, y int) {
				area++
				perimeter += 4
				sides += Corners(x, y)
				q.PushBack([2]int{x, y})
				check[x][y] = true // mark node as visited
			}

			Get := func() (int, int) {
				Node := q.Front().Value.([2]int)
				q.Remove(q.Front())
				return Node[0], Node[1]
			}

			Add(i, j)
			for q.Len() > 0 {
				x, y := Get()

				for di := 0; di < 4; di++ {
					xn, yn := x+directions[di][0], y+directions[di][1]
					if xn < 0 || xn >= n || yn < 0 || yn >= m || mat[xn][yn] != mat[x][y] {
						continue
					}

					perimeter--
					if !check[xn][yn] {
						Add(xn, yn)
					}
				}
			}

			ans += area * sides
		}
	}

	fPrintln(ans)
}
