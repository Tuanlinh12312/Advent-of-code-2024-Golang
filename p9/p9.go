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

	var s string 
	fScan(&s)

	n := len(s)
	a := make([]int, 0)
	blank, block := make([][2]int, 0), make([][2]int, 0)

	Push := func(r int, num int) {
		if r == -1 {
			blank = append(blank, [2]int{num, len(a)})
		} else {
			block = append(block, [2]int{num, len(a)})
		}
		for i := 0; i < num; i ++ {
			a = append(a, r)
		}
	}

	
	for i, cr := 0, 0; i < n; i += 2 {
		Push(cr, int(s[i]) - '0')
		if i + 1 < n {
			Push(-1, int(s[i+1]) - '0')
		}
		cr ++
	}
	
	for i := len(block) - 1; i >= 0; i -- {
		ol, op := block[i][0], block[i][1]
		for j := 0; j < len(blank); j ++ {
			al, ap := blank[j][0], blank[j][1]
			if ol <= al && op > ap {
				for k := ap; k < ap + ol; k ++ {
					a[k] = a[op]
				}
				for k := op; k < op + ol; k ++ {
					a[k] = -1
				}
				blank[j] = [2]int{al - ol, ap + ol}
				break
			}
		}
	}
	
	// for _, num := range a {
	// 	if num == -1 {
	// 		fPrintf(".")
	// 	} else {
	// 		fPrintf("%v", num)
	// 	}
	// }
	// fPrintln()
	
	ans := 0
	for i, val := range a {
		if val != -1 {
			ans += i*val
		}
	}

	fPrintln(ans)
}
