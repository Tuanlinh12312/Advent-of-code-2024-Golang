package main

// import (
// 	// "container/heap"
// 	// "container/list"

// 	// "slices"
// 	// "sort"

// 	"bufio"
// 	"fmt"
// 	"os"
// 	// "regexp"

// 	// "io"

// 	"strconv"
// 	"strings"
// 	// "regexp"
// 	// "math"
// 	// "cmp"
// )

// var (
// 	w *bufio.Writer
// 	r *bufio.Reader
// )

// func fScan(a ...any) (int, error) {
// 	return fmt.Fscan(r, a...)
// }

// func fScanf(format string, a ...any) (int, error) {
// 	return fmt.Fscanf(r, format, a...)
// }

// func fScanln(a ...any) (int, error) {
// 	return fmt.Fscanln(r, a...)
// }

// func fReadInts() ([]int, error) {
// 	inp, err := r.ReadString('\n')
// 	if err != nil {
// 		return nil, err
// 	}

// 	ans := make([]int, 0)
// 	for _, num := range strings.Fields(inp) {
// 		n, _ := strconv.Atoi(num)
// 		ans = append(ans, n)
// 	}

// 	return ans, nil
// }

// func fPrintf(format string, a ...any) (int, error) {
// 	return fmt.Fprintf(w, format, a...)
// }

// func fPrintln(a ...any) (int, error) {
// 	return fmt.Fprintln(w, a...)
// }

// func main() {
// 	// r = bufio.NewReader(os.Stdin)
// 	// w = bufio.NewWriter(os.Stdout)

// 	fin, _ := os.Open("input.txt")
// 	defer fin.Close()
// 	r = bufio.NewReader(fin)

// 	fout, _ := os.Create("output.txt")
// 	defer fout.Close()
// 	w = bufio.NewWriter(fout)

// 	defer w.Flush()

// 	a := make([]string, 0)

// 	for {
// 		var s string 
// 		_, err := fScanln(&s)
// 		if err != nil {
// 			break
// 		}

// 		a = append(a, s)
// 	}

// 	ans := 0
// 	pattern := "XMAS"
// 	n, m := len(a), len(a[0])
// 	directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

// 	count := func(i, j, dx, dy int) int {
// 		for k := 0; k < 4; k ++ {
// 			in, jn := i + dx*k, j + dy*k
// 			if in < 0 || in >= n || jn < 0 || jn >= m {
// 				return 0
// 			}
// 			if a[in][jn] != pattern[k] {
// 				return 0
// 			}
// 		}
// 		return 1
// 	}


// 	for i := 0; i < n; i ++ {
// 		for j := 0; j < m; j ++ {
// 			for k := 0; k < 8; k ++ {
// 				ans += count(i, j, directions[k][0], directions[k][1])
// 			}
// 		}
// 	}
	
// 	fPrintln(ans)
// }
