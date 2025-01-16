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

	ans := 0 
	for {
		var s string
		err := fGetln(&s)
		if err != nil {
			break
		}


		a := strings.Split(s, ":")
		sum, _ := strconv.Atoi(a[0])
		nums, _ := SplitInts(a[1], ' ')
		
		mp := make(map[int]bool)
		mp[0] = true

		app := func(a, b int) int {
			sa, sb := strconv.Itoa(a), strconv.Itoa(b)
			num, _ := strconv.Atoi(sa + sb) 
			return num
		}

		for _, num := range nums {
			nmp := make(map[int]bool)

			for crr, _ := range mp {
				nmp[crr + num] = true
				nmp[crr * num] = true
				nmp[app(crr, num)] = true
			}

			mp = nmp
		}

		if mp[sum] {
			ans += sum
		}
	}

	fPrintln(ans)
}
