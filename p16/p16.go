package main

import (
	"container/heap"
	"math"
	"sort"

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

var directions = [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

type State struct {
	x    int
	y    int
	di   int
	cost int
}

func (s State) Hash() [3]int {return [3]int{s.x, s.y, s.di}}
func (s State) HashCost() [4]int {return [4]int{s.x, s.y, s.di, s.cost}}
func (s State) HashNode() [2]int {return [2]int{s.x, s.y}}

func (s State) Move(dn int) State {
	sn := State{
		x:    s.x + directions[dn][0],
		y:    s.y + directions[dn][1],
		di:   dn,
		cost: s.cost + 1,
	}
	if dn != s.di {
		sn.cost += 1000
	}
	return sn
}

type PriorityQueue []State

func (q PriorityQueue) Len() int           { return len(q) }
func (q PriorityQueue) Less(i, j int) bool { return q[i].cost < q[j].cost }
func (q PriorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *PriorityQueue) Push(x any)        { *q = append(*q, x.(State)) }

func (q *PriorityQueue) Pop() any {
	res := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return res
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

	mat := make([][]rune, 0)
	for {
		var s string
		_, err := fScan(&s)
		if err != nil {
			break
		}
		mat = append(mat, []rune(s))
	}

	var sx, sy, ex, ey int
	n, m := len(mat), len(mat[0])

	for i := range n {
		for j := range m {
			if mat[i][j] == 'S' {
				sx, sy = i, j
			}
			if mat[i][j] == 'E' {
				ex, ey = i, j
			}
		}
	}

	cost := make(map[[3]int]int)
	q := PriorityQueue{State{
		x:    sx,
		y:    sy,
		di:   3,
		cost: 0,
	}}
	heap.Init(&q)

	for q.Len() > 0 {
		s := q.Pop().(State)
		if s.cost != cost[s.Hash()] {
			continue
		}

		for k := range 4 {
			sn := s.Move(k)
			if sn.x < 0 || sn.x >= n || sn.y < 0 || sn.y >= m || mat[sn.x][sn.y] == '#' {
				continue
			}
			if sn.cost < cost[sn.Hash()] || cost[sn.Hash()] == 0 {
				cost[sn.Hash()] = sn.cost
				q.Push(sn)
			}
		}
	}

	minAns := math.MaxInt32
	for k := range 4 {
		s := State{
			x : ex,
			y : ey,
			di : k,
		}
		if cost[s.Hash()] > 0 && cost[s.Hash()] < minAns {
			minAns = cost[s.Hash()]
		}
	}

	states := make([]State, 0)
	for key, value := range cost {
		states = append(states, State{
			x:    key[0],
			y:    key[1],
			di:   key[2],
			cost: value,
		})
	}
	sort.Slice(states, func(i, j int) bool {
		return states[i].cost > states[j].cost
	})

	IsBest := make(map[[4]int]bool)
	IsBestNode := make(map[[2]int]bool)

	MarkBest := func(s State) {
		IsBest[s.HashCost()] = true 
		IsBestNode[s.HashNode()] = true 
	}

	for _, s := range states {
		if s.x == ex && s.y == ey && s.cost == minAns {
			MarkBest(s)
		} else {
			for k := range 4 {
				if IsBest[s.Move(k).HashCost()] {
					MarkBest(s)
				}
			}
		}
	}

	fPrintln(len(IsBestNode)+1)
}
