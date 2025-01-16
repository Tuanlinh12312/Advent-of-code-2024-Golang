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
	cnt := make(map[[4]int]int)

	for {
		var num int 
		_, err := fScan(&num)
		if err != nil {
			break 
		}

		seen := make(map[[4]int]bool)
		changes := make([]int, 0)
		nums := make([]int, 0)
		nums = append(nums, num%10)

		for range 2000 {
			numOld := num

			num ^= num*64
			num %= 16777216
			num ^= num/32
			num %= 16777216
			num ^= num*2048
			num %= 16777216

			changes = append(changes, (numOld%10)-(num%10))
			nums = append(nums, num%10)
		}

		for i := 0; i+4 < len(changes); i++{
			if !seen[[4]int(changes[i:i+4])] {
				seen[[4]int(changes[i:i+4])] = true 
				cnt[[4]int(changes[i:i+4])] += nums[i+4]
			}
		}
	}

	for _, value := range cnt {
		ans = max(ans, value)
	}
	fPrintln(ans)
}
