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

func fReadInts() ([]int, error) {
	inp, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	ans := make([]int, 0)
	for _, num := range strings.Fields(inp) {
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

	ans := 0
	enabled := true
	number := regexp.MustCompile(`[0-9]{1,3}`)
	pattern := regexp.MustCompile(`(mul\(([0-9]{1,3})\,([0-9]{1,3})\))|(do\(\))|don't\(\)`)

	for {
		var s string
		_, err := fScan(&s)
		if err != nil {
			break
		}

		a := pattern.FindAllString(s, -1)

		for _, exp := range a {
			if exp == "don't()" {
				enabled = false
			} else if exp == "do()" {
				enabled = true
			} else if enabled {
				nums := number.FindAllString(exp, 2)
				e1, _ := strconv.Atoi(nums[0])
				e2, _ := strconv.Atoi(nums[1])
				ans += e1 * e2
			}
		}
	}

	fPrintln(ans)
}
