package pkg

import "math"

func GetZeroIntMatrix(n, m int) [][]int {
	matrix := make([][]int, n, n)
	for i, _ := range matrix {
		matrix[i] = make([]int, m, m)
	}

	return matrix
}

func MultMatrix(matrix [][]float64, mult float64) [][]float64 {
	newMatrix := make([][]float64, len(matrix), len(matrix))
	for i, arr := range matrix {
		newMatrix[i] = make([]float64, len(arr), len(arr))
		for j, _ := range arr {
			newMatrix[i][j] = matrix[i][j] * mult
		}
	}

	return newMatrix
}

func PowMatrix(matrix [][]float64, pow int) [][]float64 {
	newMatrix := make([][]float64, len(matrix), len(matrix))
	for i, arr := range matrix {
		newMatrix[i] = make([]float64, len(arr), len(arr))
		for j, _ := range arr {
			newMatrix[i][j] = math.Pow(matrix[i][j], float64(pow))
		}
	}

	return newMatrix
}

func PowArray(arr []float64, pow float64) []float64 {
	newArray := make([]float64, len(arr), len(arr))
	for i, v := range arr {
		newArray[i] = math.Pow(v, pow)
	}

	return newArray
}

func MultArrays(arr1, arr2 []float64) []float64 {
	resArr := make([]float64, len(arr1), len(arr1))
	for i, _ := range arr1 {
		resArr[i] = arr1[i] * arr2[i]
	}

	return resArr
}

func DivideArray(arr []float64, div float64) []float64 {
	resArr := make([]float64, len(arr), len(arr))
	for i, _ := range arr {
		resArr[i] = arr[i] / div
	}

	return resArr
}

func ArraySum(arr []float64) float64 {
	sum := 0.0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func ArrayAvg(arr []int) int {
	sum := 0
	count := 0
	for _, v := range arr {
		sum += v
		count += 1
	}

	return sum / count
}

func ArrayFloat64Avg(arr []float64) float64 {
	sum := 0.0
	count := 0.0
	for _, v := range arr {
		sum += v
		count += 1.0
	}

	return sum / count
}
