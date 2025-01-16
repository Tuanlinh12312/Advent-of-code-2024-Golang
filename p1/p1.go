
package main
 
import (
    // "container/heap"
    // "container/list"
 
    // "slices"
    // "sort"
 
    "bufio"
    "fmt"
    "os"
	// "io"
 
    // "strings"
    // "strconv"
    // "regexp"
    // "math"
    // "cmp"
)

var (
    w *bufio.Writer
    r *bufio.Reader
)

func fscan(a ...any) (int, error) {
	return fmt.Fscan(r, a...)
}

func fscanf(format string, a ...interface{}) (int, error) {
	return fmt.Fscanf(r, format, a...)
}

func fprintf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(w, format, a...)
}

func main(){
	// r = bufio.NewReader(os.Stdin)
    // w = bufio.NewWriter(os.Stdout)
	
	fin, _ := os.Open("input.txt")
    defer fin.Close()
    r = bufio.NewReader(fin)

	fout, _ := os.Create("output.txt")
	defer fout.Close()
	w = bufio.NewWriter(fout)
	defer w.Flush()

	A := make([]int, 0)
	cnt := make(map[int]int)

	for {
		var a, b int
		_, err := fscan(&a, &b)
		if err != nil {
			break;
		}
		
		A = append(A, a)
		cnt[b] ++
	}

	ans := 0
	for _, x := range A {
		ans += x * cnt[x]
	}

	fprintf("%d\n", ans)
}