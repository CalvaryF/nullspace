package nullspace


import (
	"fmt"
)


// Performs matrix multiplication
// The left matrix transforms each column vector in the right matrix
// Outputting a new basis / transform combining the two
func matrixMultiplication(a, b [][]int) ([][]int, error) {
	// Get dimensions
	rowsA := len(a)
	colsA := len(a[0])
	rowsB := len(b)
	colsB := len(b[0])

	// Check if matrices are compatible for multiplication
	if colsA != rowsB {
		return nil, fmt.Errorf("matrices cannot be multiplied: incompatible dimensions")
	}

	// Initialize the result matrix
	result := MakeMatrix(rowsA, colsB)

	// Perform matrix multiplication
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return result, nil
}

