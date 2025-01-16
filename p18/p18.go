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

	bytes := make([][]int, 0)
	for {
		nums, err := fReadInts()
		if err != nil {
			break
		}
		bytes = append(bytes, nums)
	}

	Check := func(mid int) bool {
		directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
		mat := [71][71]int{}
		dis := [71][71]int{}
	
		for i := range mid {
			mat[bytes[i][0]][bytes[i][1]] = 1
		}
	
		q := list.New()
		q.PushBack([2]int{0, 0})
		dis[0][0] = 1
	
		for q.Len() > 0 {
			Node := q.Front().Value.([2]int)
			q.Remove(q.Front())
			x, y := Node[0], Node[1]
			
			for k := range 4 {
				xn, yn := x+directions[k][0], y+directions[k][1]
				if xn < 0 || xn >= 71 || yn < 0 || yn >= 71 || mat[xn][yn] == 1 {
					continue
				}
				if dis[xn][yn] == 0 {
					dis[xn][yn] = dis[x][y] + 1
					q.PushBack([2]int{xn, yn})
				}
			}
		}

		return dis[70][70] != 0
	}

	lo, hi := 1024, len(bytes)
	for lo < hi {
		mid := (lo+hi+1)/2
		if Check(mid) {
			lo = mid
		} else {
			hi = mid-1
		}
	}
	fPrintf("%v,%v\n", bytes[lo][0], bytes[lo][1])
}
