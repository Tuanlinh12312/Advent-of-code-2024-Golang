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

	ans := 0
	locks := make([][5]int, 0)
	keys := make([][5]int, 0)

	for {
		mat := make([][]rune, 7)
		err := error(nil)
		for i := range 7 {
			var s string
			_, er := fScan(&s)

			if er != nil {
				err = er
			} else {
				mat[i] = []rune(s)
			}
		}
		if err != nil {
			break
		}

		if mat[0][0] == '#' {
			lock := [5]int{}

			for i := range 5 {
				for j := range 7 {
					if mat[j][i] == '#' {
						lock[i] = j
					}
				}
			}

			locks = append(locks, lock)
		} else {
			key := [5]int{}

			for i := range 5 {
				for j := range 7 {
					if mat[6-j][i] == '#' {
						key[i] = j
					}
				}
			}

			keys = append(keys, key)
		}
	}

	for _, lock := range locks {
		for _, key := range keys {
			fit := true 
			for i := range 5 {
				if lock[i] + key[i] > 5 {
					fit = false
				}
			}
			if fit {
				ans++
			}
		}
	}
	fPrintln(ans)
}
