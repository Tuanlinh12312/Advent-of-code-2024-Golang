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
	for _, num := range strings.FieldsFunc(inp, func(r rune) bool{
		return r == sep
	}){
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

	A := make(map[[2]int]bool)

	for {
		a, _ := fReadInts('|')
		if len(a) == 0 {
			break
		}
		
		A[[2]int{a[0], a[1]}] = true
	}

	CheckOrder := func(a []int) bool {
		n := len(a)
		for i := 0; i < n; i ++ {
			for j := i + 1; j < n; j ++ {
				if A[[2]int{a[j], a[i]}] {
					return false
				}
			}
		}
		return true
	}

	ans := 0
	for {
		a, err := fReadInts(',')
		if err != nil {
			break
		}

		if !CheckOrder(a) {
			res := make([]int, 0)
			check := make(map[int]bool)

			for len(res) < len(a) {
				for _, num := range a {
					if check[num] {
						continue
					}
	
					ok := true
					for _, other := range a {
						if !check[other] && A[[2]int{other, num}] {
							ok = false
						}
					}
	
					if ok {
						res = append(res, num)
						check[num] = true
					}
				}
			}

			ans += res[len(a)/2]
		}
	}
	
	fPrintln(ans)
}
