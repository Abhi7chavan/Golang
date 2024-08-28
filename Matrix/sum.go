package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var t int
	fmt.Scan(&t)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < t; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sizes := strings.Split(line, " ")
		M, N := sizes[0], sizes[1]
		integerM, _ := strconv.Atoi(M)
		integerN, _ := strconv.Atoi(N)
		matrix := ReadMatrix(integerM, integerN, reader)
		sum := CalculateSum(integerM, integerN, matrix, 0, 0)
		fmt.Println("Sum of matrix elements:", sum)
	}
}

func ReadMatrix(M int, N int, reader *bufio.Reader) [][]int {
	result := make([][]int, M)
	for i := 0; i < M; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sizes := strings.Split(line, " ")
		var subarray []int
		for _, value := range sizes {
			valueInt, _ := strconv.Atoi(value)
			subarray = append(subarray, valueInt)
		}
		result[i] = subarray
	}
	return result
}

func CalculateSum(M, N int, matrix [][]int, i, j int) int {
	if i == M {
		return 0;
	}
	if j == N {
		return CalculateSum(M, N, matrix, i+1, 0)
	}

	return matrix[i][j] + CalculateSum(M, N, matrix, i, j+1)
}
