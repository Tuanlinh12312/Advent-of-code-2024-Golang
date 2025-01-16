package main

import (
	// "container/heap"
	// "container/list"

	// "slices"
	// "sort"

	"bufio"
	"fmt"
	"os"
	"sort"

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

	isNode := make(map[string]bool)
	adj := make(map[string][]string)
	connected := make(map[[2]string]bool)

	for {
		edge, err := r.ReadString('\n')
		if err != nil {
			break
		}

		character := regexp.MustCompile(`[a-z]+`)
		nodes := character.FindAllString(edge, -1)

		isNode[nodes[0]] = true
		isNode[nodes[1]] = true
		adj[nodes[0]] = append(adj[nodes[0]], nodes[1])
		adj[nodes[1]] = append(adj[nodes[1]], nodes[0])
		connected[[2]string{nodes[0], nodes[1]}] = true
		connected[[2]string{nodes[1], nodes[0]}] = true
	}

	ans := 0
	ansClique := ""

	CheckClique := func(clique []string) {
		for _, u := range clique {
			for _, v := range clique {
				if u != v && !connected[[2]string{u, v}] {
					return
				}
			}
		}
		if ans < len(clique) {
			ans = len(clique)
			sort.Strings(clique)
			ansClique = ""
			for _, s := range clique {
				if len(ansClique) == 0 {
					ansClique += s
				} else {
					ansClique += "," + s
				}
			}
		}
	}

	for A := range isNode {
		for mask := 0; mask < 1<<len(adj[A]); mask++ {
			clique := []string{A}
			for i := range len(adj[A]) {
				if mask>>i&1 != 0 {
					clique = append(clique, adj[A][i])
				}
			}
			CheckClique(clique)
		}
	}

	fPrintln(ansClique)
}
