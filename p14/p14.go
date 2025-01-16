package main

import (
	// "container/heap"
	// "container/list"

	// "slices"
	"sort"

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

	number := regexp.MustCompile(`(\d|-)+`)
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

	n, m := 101, 103
	robots := make([][]int, 0)

	for {
		nums, err := fReadInts()
		if err != nil {
			break
		}

		robots = append(robots, nums)
	}
	
	Display := func(seconds int) {
		check := make(map[[2]int]bool)

		for _, robot := range robots {
			px, py, vx, vy := robot[0], robot[1], robot[2], robot[3]
			px = ((px + vx*seconds)%n + n) % n
			py = ((py + vy*seconds)%m + m) % m
			check[[2]int{px, py}] = true
		}

		for i := range n {
			for j := range m {
				if check[[2]int{i, j}] {
					fPrintf("x")
				} else {
					fPrintf(".")
				}
			}
			fPrintln()
		}
	}

	LongestLine := func(seconds int) int {
		check := make(map[[2]int]bool)

		for _, robot := range robots {
			px, py, vx, vy := robot[0], robot[1], robot[2], robot[3]
			px = ((px + vx*seconds)%n + n) % n
			py = ((py + vy*seconds)%m + m) % m
			check[[2]int{px, py}] = true
		}

		ans, cr := 0, 0
		for i := range n {
			for j := range m {
				if check[[2]int{i, j}] {
					cr ++ 
					ans = max(ans, cr)
				} else {
					cr = 0
				}
			}
		}

		return ans
	}

	rank := make([][2]int, 0)
	for seconds := range n*m {
		rank = append(rank, [2]int{seconds, LongestLine(seconds)})
	}
	sort.Slice(rank, func(i, j int) bool {
		return rank[i][1] > rank[j][1]
	})

	for i := range 5 {
		fPrintf("second %v:\n", rank[i][0])
		Display(rank[i][0])
		fPrintln()
	}
}
