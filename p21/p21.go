package main

import (
	// "container/heap"
	// "container/list"

	// "slices"
	// "sort"

	"bufio"
	"fmt"
	"math"
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

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

	nums := "1234567890A"
	numpad := []string{"789", "456", "123", "X0A"}
	numpadPosition := make(map[rune][2]int)
	for i := range numpad {
		for j, c := range numpad[i] {
			numpadPosition[c] = [2]int{i, j}
		}
	}

	moves := "<v>^A"
	movepad := []string{"X^A", "<v>"}
	movepadPosition := make(map[rune][2]int)
	for i := range movepad {
		for j, c := range movepad[i] {
			movepadPosition[c] = [2]int{i, j}
		}
	}

	steps := 26 // number of directional pads
	dp := make([]map[rune]map[rune]int, steps)
	for i := range steps {
		dp[i] = make(map[rune]map[rune]int)
	}

	CalcSequence := func(t int, s string) int {
		ans := 0
		s = "A" + s + "A"
		for i := 0; i+1 < len(s); i++ {
			ans += dp[t-1][rune(s[i])][rune(s[i+1])]
		}
		return ans
	}
	MakeSequence := func(num int, r rune) string {
		ans := ""
		for range num {
			ans += string(r)
		}
		return ans
	}

	for i := range steps {
		var isNum bool
		var characters string
		var Position map[rune][2]int

		if i < steps-1 {
			isNum = false
			characters = moves
			Position = movepadPosition
		} else {
			isNum = true
			characters = nums
			Position = numpadPosition
		}

		for _, r1 := range characters {
			x1, y1 := Position[r1][0], Position[r1][1]
			dp[i][r1] = make(map[rune]int)

			for _, r2 := range characters {
				x2, y2 := Position[r2][0], Position[r2][1]

				dp[i][r1][r2] = math.MaxInt32
				Update := func(cmp int) {
					dp[i][r1][r2] = min(dp[i][r1][r2], cmp)
				}

				if i == 0 {
					Update(Abs(x1-x2) + Abs(y1-y2) + 1)
				} else {
					dp[i][r1][r2] = math.MaxInt64
					if x1 < x2 && y1 < y2 {
						if !isNum || y1 != 0 || x2 != 3 {
							Update(CalcSequence(i, MakeSequence(x2-x1, 'v')+MakeSequence(y2-y1, '>')))
						}
						Update(CalcSequence(i, MakeSequence(y2-y1, '>')+MakeSequence(x2-x1, 'v')))

					} else if x1 >= x2 && y1 < y2 {
						if isNum || y1 != 0 || x2 != 0 {
							Update(CalcSequence(i, MakeSequence(x1-x2, '^')+MakeSequence(y2-y1, '>')))
						}
						Update(CalcSequence(i, MakeSequence(y2-y1, '>')+MakeSequence(x1-x2, '^')))
					} else if x1 < x2 && y1 >= y2 {
						if isNum || x1 != 0 || y2 != 0 {
							Update(CalcSequence(i, MakeSequence(y1-y2, '<')+MakeSequence(x2-x1, 'v')))
						}
						Update(CalcSequence(i, MakeSequence(x2-x1, 'v')+MakeSequence(y1-y2, '<')))
					} else {
						if !isNum || x1 != 3 || y2 != 0 {
							Update(CalcSequence(i, MakeSequence(y1-y2, '<')+MakeSequence(x1-x2, '^')))
						}
						Update(CalcSequence(i, MakeSequence(x1-x2, '^')+MakeSequence(y1-y2, '<')))
					}
				}
			}
		}
	}

	ans := 0
	for {
		var s string
		_, err := fScan(&s)
		if err != nil {
			break
		}

		minLength := CalcSequence(steps, s[:3])
		number, _ := strconv.Atoi(s[:3])
		ans += minLength * number
	}

	fPrintln(ans)
}
