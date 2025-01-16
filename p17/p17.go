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

	fReadInts()
	fReadInts()
	fReadInts()
	fReadInts()
	instructions, _ := fReadInts()

	Try := func(A int64) []int {
		B, C := int64(0), int64(0)
		output := make([]int, 0)

		Combo := func(operand int64) int64 {
			if operand == 4 {
				return A
			} else if operand == 5 {
				return B
			} else if operand == 6 {
				return C
			} else {
				return operand
			}
		}

		for i := int64(0); int(i) < len(instructions); i += 2 {
			opcode, operand := int64(instructions[i]), int64(instructions[i+1])

			if opcode == 0 {
				A /= (1 << Combo(operand))
			} else if opcode == 1 {
				B ^= operand
			} else if opcode == 2 {
				B = Combo(operand) % 8
			} else if opcode == 3 {
				if A != 0 {
					i = operand - int64(2)
				}
			} else if opcode == 4 {
				B ^= C
			} else if opcode == 5 {
				output = append(output, int(Combo(operand)%int64(8)))
			} else if opcode == 6 {
				B = A / (1 << Combo(operand))
			} else if opcode == 7 {
				C = A / (1 << Combo(operand))
			}
		}

		return output
	}

	A := [48]int{}
	check := [48]bool{}

	Bit := func(num, b int) int {
		return num >> b & 1
	}

	CheckWrite := func(pos, num int) bool {
		for i := range 3 {
			if pos+i >= 48 {
				if Bit(num, i) == 1 {
					return false
				}
			} else if check[pos+i] && A[pos+i] != Bit(num, i) {
				return false
			}
		}
		return true
	}

	Write := func(pos, num int) {
		for i := range 3 {
			if pos+i < 48 {
				A[pos+i] = Bit(num, i)
				check[pos+i] = true
			}
		}
	}

	var BackTrack func(int)
	BackTrack = func(i int) {
		if i == len(instructions) {
			ans := int64(0)
			for i := range len(A) {
				ans ^= int64(A[i]) << i
			}
			fPrintln(ans)
			fPrintln(Try(ans))
			return
		}

		target := instructions[i]

		for B := range 8 {
			if !CheckWrite(i*3, B) {
				continue
			}
			AOld, checkOld := A, check
			Write(i*3, B)
			if CheckWrite(i*3+(B^7), target^B) {
				Write(i*3+(B^7), target^B)
				BackTrack(i + 1)
			}
			A, check = AOld, checkOld
		}
	}

	BackTrack(0)
}
