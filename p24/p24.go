package main

import (
	// "container/heap"
	// "container/list"

	// "slices"
	"sort"

	"bufio"
	"fmt"
	"os"

	"regexp"

	// "io"

	"strconv"
	// "strings"
	// "regexp"
	// "math"
	// "cmp"

	// "math/rand"
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

	bits := 45
	gate := make(map[string][]string)
	dependent := make(map[string][]string)
	checkGate := make([]map[string]bool, bits+1)

	for {
		var num int
		var s string
		n, _ := fScanln(&s, &num)
		if n < 2 {
			break
		}
	}

	for {
		var inp1, op, inp2, out string
		_, err := fScanf("%s %s %s -> %s\n", &inp1, &op, &inp2, &out)
		if err != nil {
			break
		}

		dependent[inp1] = append(dependent[inp1], out)
		dependent[inp2] = append(dependent[inp2], out)
		gate[out] = append(gate[out], inp1, op, inp2)
	}

	var DfsCheckGate func(b int, g string) 
	DfsCheckGate = func(b int, g string) {
		checkGate[b][g] = true 
		if gate[g] != nil {
			if !checkGate[b][gate[g][0]] {
				DfsCheckGate(b, gate[g][0])
			}
			if !checkGate[b][gate[g][2]] {
				DfsCheckGate(b, gate[g][2])
			}
		}
	}

	InitCheck := func(){
		checkGate = make([]map[string]bool, bits+1)
		for i := range bits+1 {
			checkGate[i] = make(map[string]bool)
		}
		for i := range bits+1 {
			gateZi := "z" + fmt.Sprintf("%02d", i)
			DfsCheckGate(i, gateZi)
		}
	}
	
	InitCheck()
	FixName := func(s string) string {
		if gate[s] != nil {
			return s + "_" + gate[s][1]
		}
		return s
	}

	for g := range gate {
		for i := range 10 {
			if checkGate[i][g] {
				fPrintln(FixName(g) , FixName(gate[g][0]))
				fPrintln(FixName(g), FixName(gate[g][2]))
				break
			}
		}
	}

	sus := make(map[string]bool)
	for g := range gate {
		if gate[g][1] != "XOR" && g[0] == 'z' && g[1:] != "45" {
			fPrintf("sus: %s is a zXX gate but is not XOR\n", g)
			sus[g] = true
		}
		if gate[g][1] == "XOR" && gate[g][0][0] != 'x' && gate[g][0][0] != 'y' && g[0] != 'z' {
			fPrintf("sus: %s is a XOR gate but is not connected to x, y or z\n", g)
			sus[g] = true
		}
		if gate[g][1] == "OR" && gate[gate[g][0]][1] != "AND" {
			fPrintf("sus: %s has parent OR but not an AND gate\n", gate[g][0])
			sus[gate[g][0]] = true
		}
		if gate[g][1] == "OR" && gate[gate[g][2]][1] != "AND" {
			fPrintf("sus: %s has parent OR but not an AND gate\n", gate[g][2])
			sus[gate[g][2]] = true
		}
		if gate[g][1] == "AND" && len(dependent[g]) == 2 && gate[g][0][1:] != "00" {
			fPrintf("sus: %s is AND gate but has != 1 dependencies\n", g)
			sus[g] = true
		}
	}

	ans := make([]string, 0) 
	for susgate := range sus {
		ans = append(ans, susgate)
	}
	sort.Strings(ans)
	for i, s := range ans {
		if i == 0 {
			fPrintf("%s", s)
		} else {
			fPrintf(",%s", s)
		}
	}
}
