package main

import (
	// "container/heap"
	"container/list"

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

func Abs(a int) int {
	if a < 0 {
		return -a 
	}
	return a
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
		mat = append(mat, []rune(s))
	}

	n, m := len(mat), len(mat[0])
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
	
	CalculateDis := func(X, Y int) [][]int {
		dis := make([][]int, n)
		for i := range n {
			dis[i] = make([]int, m)
			for j := range m {
				dis[i][j] = -1
			}
		}
		
		q := list.New()
		q.PushBack([2]int{X, Y})
		dis[X][Y] = 0
	
		for q.Len() > 0 {
			Node := q.Front().Value.([2]int)
			q.Remove(q.Front())
			x, y := Node[0], Node[1]
			
			for k := range 4 {
				xn, yn := x+directions[k][0], y+directions[k][1]
				if xn < 0 || xn >= n || yn < 0 || yn >= m || mat[xn][yn] == '#' {
					continue
				}
				if dis[xn][yn] == -1 {
					dis[xn][yn] = dis[x][y] + 1
					q.PushBack([2]int{xn, yn})
				}
			}
		}

		return dis
	}

	var sx, sy, ex, ey int 
	for i := range n {
		for j := range m {
			if mat[i][j] == 'S' {
				sx, sy = i, j
			}
			if mat[i][j] == 'E' {
				ex, ey = i, j 
			}
		}
	}
	ans := 0
	Sdis := CalculateDis(sx, sy) 
	Edis := CalculateDis(ex, ey) 

	for si := range n {
		for sj := range m {
			for ei := range n {
				for ej := range m {
					if Abs(si-ei) + Abs(sj-ej) <= 20 {
						if Sdis[si][sj] + Edis[ei][ej] + 100 + Abs(si-ei) + Abs(sj-ej) <= Sdis[ex][ey]  && Sdis[si][sj] != -1 && Edis[ei][ej] != -1 {
							ans ++
							// fPrintf("start: (%v, %v), end: (%v, %v)\n", si, sj, ei, ej)
							// fPrintf("Save: %v\n", Sdis[ex][ey] - Sdis[si][sj] - Edis[ei][ej])
						}
					}
				}
			}
		}
	}
	fPrintln(ans)
}
