package main

import (
	"strconv"
	"bufio"
	"fmt"
	"os"
)

func mul(a [4]int, b[4]int) [4]int {
	ret := [4]int{}
	ret[0] = a[0]*b[0] + a[1]*b[2]
	ret[1] = a[0]*b[1] + a[1]*b[3]
	ret[2] = a[2]*b[0] + a[3]*b[2]
	ret[3] = a[2]*b[1] + a[3]*b[3]
	return ret
}

func fib(n int) int {
	if n <= 1 { return n }
	result := [4]int{1, 0, 0, 1}
	matrix := [4]int{1, 1, 1, 0}

	for {
		if n < 1 { break; }
		if n % 2 == 1 {
			result = mul(matrix, result)
		}

		matrix = mul(matrix, matrix)
		n /= 2
	}

	return result[2]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l, _ := strconv.Atoi(scanner.Text())
	i := 0
	for {
		if i >= l { break }
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		fmt.Println(fib(n))
		i += 1
	}
}
