package nullspace


import (
	"fmt"
)

//Swaps the row and colum indices of matrix a 
func Transpose(a [][]int)([][]int, error){
    
    if err := checkMatrixNotNil(a); err != nil {
        return nil, err
    }
    if err := checkMatrixNotEmpty(a); err != nil {
        return nil, err
    }
    if err := checkMatrixRectangular(a); err != nil {
        return nil, err
    }


	rows := len(a)
	cols := len(a[0])
	result := MakeMatrix(cols, rows)

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
           result[j][i] = a[i][j]
        } 
    }
    return result, nil
}

// Performs matrix multiplication
// The left matrix transforms each column vector in the right matrix
// Outputting a new basis / transform combining the two
func MatrixMultiplication(a, b [][]int) ([][]int, error) {

     if err := checkMatrixNotNil(a); err != nil {
        return nil, err
    }
    if err := checkMatrixNotNil(b); err != nil {
        return nil, err
    }
    if err := checkMatrixNotEmpty(a); err != nil {
        return nil, err
    }
    if err := checkMatrixNotEmpty(b); err != nil {
        return nil, err
    }
    if err := checkMatrixRectangular(a); err != nil {
        return nil, err
    }
    if err := checkMatrixRectangular(b); err != nil {
        return nil, err
    }
    if err := checkMatrixMultiplicationDims(a, b); err != nil {
        return nil, err
    }


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


