package nullspace

import (
	"fmt"
)

// Initializes a matrix of a given size
func MakeMatrix(rows, cols int) [][]int {
	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
	}
	return result
}


// Helper to print a matrix
func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}

//Creates a triangular matrix (upper right is all 1s)
func CreateTriangular(dim int)[][]int{
    result := MakeMatrix(dim, dim)
    for i := 0; i < dim; i++ {
        for j :=0; j< dim; j++{
            if(i<j){
                result[i][j] = 1

            }
        }
    } 
    return result
}
