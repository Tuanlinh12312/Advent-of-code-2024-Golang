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
	if len(inp) == 0 {
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

	a, _ := fReadInts(' ')
	cnt := make(map[[2]int]int, 0)

	var Calc func(int, int) int 
	Calc = func(num, t int) int {
		if cnt[[2]int{num, t}] > 0 {
			return cnt[[2]int{num, t}]
		}
		if t == 0 {
			return 1
		}

		if num == 0 {
			cnt[[2]int{num, t}] = Calc(1, t-1)
		} else if len(strconv.Itoa(num))%2 == 1 {
			cnt[[2]int{num, t}] = Calc(num*2024, t-1)
		} else {
			s := strconv.Itoa(num)
			fi, _ := strconv.Atoi(s[0 : len(s)/2])
			se, _ := strconv.Atoi(s[len(s)/2:])
			cnt[[2]int{num, t}] = Calc(fi, t-1) + Calc(se, t-1)
		}
		
		return cnt[[2]int{num, t}]
	}

	ans :=0
	for _, num := range a {
		ans += Calc(num, 75)
	}
	fPrintln(ans)
}
