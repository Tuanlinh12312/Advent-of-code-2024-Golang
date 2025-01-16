
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
 
    "strings"
    "strconv"
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

func fScann(a ...any) (int, error) {
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

func AbsInt(a int) int {
	if a < 0 {
		return - a
	}
	return a
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

    CheckSafe := func(a []int) bool {
        n := len(a)
        for i := 1; i < n; i++ {
            if AbsInt(a[i] - a[i-1]) > 3 {
                return false
            }
            if i > 1 && (a[i] - a[i-1])*(a[i-1] - a[i-2]) <= 0 {
                return false
            }
        }
    
        return true
    }        

    CheckSafeReal := func(a []int) bool {
        n := len(a)
        for i := 0; i < n; i ++ {
            an := append(append([]int(nil), a[0:i]...), a[i+1:n]...)
            if CheckSafe(an) {
                return true
            }
        }
        return false
    }

    ans := 0
    for {
        a, err := fReadInts()
        if err != nil {
            break
        }

        if CheckSafeReal(a){
            ans ++
        }
    }

    fPrintln(ans)
}