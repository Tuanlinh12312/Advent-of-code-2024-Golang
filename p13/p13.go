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
	epsilon := 1e-9 // Margin of error
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

	ans := 0

	for {
		buttonA, err := fReadInts()
		buttonB, _ := fReadInts()
		prize, _ := fReadInts()
		fReadInts()
		if err != nil {
			break
		}

		Ax, Ay := float64(buttonA[0]), float64(buttonA[1])
		Bx, By := float64(buttonB[0]), float64(buttonB[1])
		px, py := float64(prize[0]), float64(prize[1])

		px += 10000000000000
		py += 10000000000000

		d := px*By - py*Bx
		m := Ax*By - Ay*Bx

		fPrintf("%v %v\n", d, m)
		if m == 0 || d*m < 0 {
			continue
		}

		A := d / m
		B := (px - A*Ax) / Bx

		if !IsInt(A) || !IsInt(B) {
			continue
		}

		ans += int(math.Round(A)*3 + math.Round(B))
	}

	fPrintln(ans)
}
