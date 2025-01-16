package main

import (
	// "container/heap"
	// "container/list"

	// "slices"
	// "sort"

	"bufio"
	"fmt"
	"os"

	"regexp"

	// "io"

	"strconv"
	"strings"

	// "regexp"
	"math"
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

func fReadInts() ([]int, error) {
	inp, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	number := regexp.MustCompile(`\d+`)
	nums := number.FindAllString(inp, -1)
	res := make([]int, 0)

	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		res = append(res, n)
	}

	return res, nil
}

func fPrintf(format string, a ...any) (int, error) {
	return fmt.Fprintf(w, format, a...)
}

func fPrintln(a ...any) (int, error) {
	return fmt.Fprintln(w, a...)
}

func IsInt(a float64) bool {
	epsilon := 1e-9
	_, frac := math.Modf(math.Abs(a))
	return frac < epsilon || frac > 1.0-epsilon
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

	mat := make([][]rune, 0)
	for {
		var s string
		fScanln(&s)
		if len(s) == 0 {
			break
		}

		row := make([]rune, 0)
		for _, c := range s {
			if c == 'O' {
				row = append(row, '[', ']')
			} else if c == '#' {
				row = append(row, c, c)
			} else {
				row = append(row, c, '.')
			}
		}
		mat = append(mat, row)
	}

	n, m := len(mat), len(mat[0])

	var x, y int
	for i := range n {
		for j := range m {
			if mat[i][j] == '@' {
				x, y = i, j
			}
		}
	}

	instructions := ""

	for {
		var s string
		_, err := fScanln(&s)
		if err != nil {
			break
		}
		instructions += s
	}

	var CheckPush func(int, int, int, int, int) bool
	CheckPush = func(row, l, r, dx, dy int) bool {
		if dy == -1 {
			if mat[row][l-1] == '.' {
				return true
			} else if mat[row][l-1] == '#' {
				return false
			} else {
				return CheckPush(row, l-2, l-1, dx, dy)
			}
		} else if dy == 1 {
			if mat[row][r+1] == '.' {
				return true
			} else if mat[row][r+1] == '#' {
				return false
			} else {
				return CheckPush(row, r+1, r+2, dx, dy)
			}
		} else {
			for i := l; i <= r; i++ {
				if mat[row+dx][i] == '#' {
					return false
				}
				if mat[row+dx][i] == '[' {
					if !CheckPush(row+dx, i, i+1, dx, dy) {
						return false
					}
				}
				if mat[row+dx][i] == ']' && i == l {
					if !CheckPush(row+dx, i-1, i, dx, dy) {
						return false
					}
				}
			}
			return true
		}
	}

	var Push func(int, int, int, int, int)
	Push = func(row, l, r, dx, dy int) {
		if dy == -1 {
			if mat[row][l-1] == ']' {
				Push(row, l-2, l-1, dx, dy)
			}
			mat[row][l-1], mat[row][l], mat[row][r] = mat[row][l], mat[row][r], mat[row][l-1]
		} else if dy == 1 {
			if mat[row][r+1] == '[' {
				Push(row, r+1, r+2, dx, dy)
			} 
			mat[row][r+1], mat[row][r], mat[row][l] = mat[row][r], mat[row][l], mat[row][r+1]
		} else {
			for i := l; i <= r; i++ {
				if mat[row+dx][i] == '[' {
					Push(row+dx, i, i+1, dx, dy)
				}
				if mat[row+dx][i] == ']' && i == l {
					Push(row+dx, i-1, i, dx, dy)
				}
			}
			for i := l; i <= r; i++ {
				mat[row][i], mat[row+dx][i] = mat[row+dx][i], mat[row][i]
			}
		}
	}

	for _, c := range instructions {
		var dx, dy int
		if c == '<' {
			dx, dy = 0, -1
		} else if c == '>' {
			dx, dy = 0, 1
		} else if c == '^' {
			dx, dy = -1, 0
		} else {
			dx, dy = 1, 0
		}

		if CheckPush(x, y, y, dx, dy) {
			Push(x, y, y, dx, dy)
			x += dx
			y += dy
		}

		// fPrintf("Move %c\n", c)
		// for i := range len(mat) {
		// 	for j := range len(mat[i]) {
		// 		fPrintf("%c", mat[i][j])
		// 	}
		// 	fPrintln()
		// }
		// fPrintln()
	}

	ans := 0
	for i := range n {
		for j := range m {
			if mat[i][j] == '[' {
				ans += 100*i + j
			}
		}
	}

	fPrintln(ans)
}
