package main

// import (
// 	// "container/heap"
// 	"container/list"

// 	// "slices"
// 	// "sort"

// 	"bufio"
// 	"fmt"
// 	"os"

// 	"regexp"

// 	// "io"

// 	"strconv"
// 	// "strings"
// 	// "regexp"
// 	// "math"
// 	// "cmp"

// 	// "math/rand"
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

// func fGetln(s *string) error {
// 	sn, err := r.ReadString('\n')
// 	if err != nil {
// 		return err
// 	}
// 	*s = sn
// 	return nil
// }

// func fReadInts() ([]int, error) {
// 	inp, err := r.ReadString('\n')
// 	if err != nil {
// 		return nil, err
// 	}

// 	number := regexp.MustCompile(`\d+`)
// 	nums := number.FindAllString(inp, -1)
// 	res := make([]int, 0)

// 	for _, num := range nums {
// 		n, _ := strconv.Atoi(num)
// 		res = append(res, n)
// 	}

// 	return res, nil
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

// 	bits := 45
// 	// // ans := make([]string, 0)

// 	bit := make(map[string]int)
// 	cntDependent := make(map[string]int)
// 	gate := make(map[string][]string)
// 	dependent := make(map[string][]string)
// 	// checkGate := make([]map[string]bool, bits+1)

// 	for {
// 		var num int
// 		var s string
// 		n, _ := fScanln(&s, &num)
// 		if n < 2 {
// 			break
// 		}

// 		bit[s[:len(s)-1]] = num
// 	}

// 	for {
// 		var inp1, op, inp2, out string
// 		_, err := fScanf("%s %s %s -> %s\n", &inp1, &op, &inp2, &out)
// 		if err != nil {
// 			break
// 		}

// 		dependent[inp1] = append(dependent[inp1], out)
// 		dependent[inp2] = append(dependent[inp2], out)
// 		cntDependent[out] = 2
// 		gate[out] = append(gate[out], inp1, op, inp2)
// 	}

// 	for out := range gate {
// 		cntDependent[out] = 2
// 	}

// 	q := list.New()
// 	for s := range bit {
// 		q.PushBack(s)
// 	}

// 	for q.Len() > 0 {
// 		s := q.Front().Value.(string)
// 		q.Remove(q.Front())

// 		for _, out := range dependent[s] {
// 			cntDependent[out]--

// 			if cntDependent[out] == 0 {
// 				inp1, op, inp2 := bit[gate[out][0]], gate[out][1], bit[gate[out][2]]
// 				if op == "AND" {
// 					bit[out] = inp1 & inp2
// 				} else if op == "OR" {
// 					bit[out] = inp1 | inp2
// 				} else {
// 					bit[out] = inp1 ^ inp2
// 				}

// 				q.PushBack(out)
// 			}
// 		}
// 	}

// 	ans := 0
// 	for i := 0; i < bits+1; i++ {
// 		gate := "z" + fmt.Sprintf("%02d", i)
// 		ans ^= bit[gate] << i
// 	}
// 	fPrintln(ans)
// }
