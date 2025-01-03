package main

import (
	"fmt"
)

func main(){ 

	
	// Print result
	fmt.Println("Result:")
    //	printMatrix(result)

    printMatrix(createTriangular(10))
}

// Initializes a matrix of a given size
func makeMatrix(rows, cols int) [][]int {
	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
	}
	return result
}

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
	result := makeMatrix(rowsA, colsB)

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

// Helper to print a matrix
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}

func createTriangular(dim int)[][]int{
    result := makeMatrix(dim, dim)
    for i := 0; i < dim; i++ {
        for j :=0; j< dim; j++{
            if(i<j){
                result[i][j] = 1

            }
        }
    } 
    return result
}

