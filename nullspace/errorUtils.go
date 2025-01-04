package nullspace

import(
    "fmt"
)

func checkMatrixNotNil(m [][]int) error {
    if m == nil {
        return fmt.Errorf("matrix is nil")
    }
    return nil
}

func checkMatrixNotEmpty(m [][]int) error {
    if len(m) == 0 || len(m[0]) == 0 {
        return fmt.Errorf("matrix is empty")
    }
    return nil
}

func checkMatrixRectangular(m [][]int) error {
    if len(m) == 0 {
        return nil // Empty matrices are trivially rectangular
    }
    cols := len(m[0])
    for i, row := range m {
        if len(row) != cols {
            return fmt.Errorf("row %d has length %d, expected %d", i, len(row), cols)
        }
    }
    return nil
}

func checkMatrixSquare(m [][]int) error {
    if len(m) != len(m[0]) {
        return fmt.Errorf("matrix must be square: %dx%d", len(m), len(m[0]))
    }
    return nil
}

func checkMatrixSymmetric(m [][]int) error {
    rows, cols := len(m), len(m[0])
    if rows != cols {
        return fmt.Errorf("matrix must be square to check symmetry")
    }
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if m[i][j] != m[j][i] {
                return fmt.Errorf("matrix is not symmetric")
            }
        }
    }
    return nil
}

func checkVectorLength(v1, v2 []int) error {
    if len(v1) != len(v2) {
        return fmt.Errorf("vectors must have the same length: %d and %d", len(v1), len(v2))
    }
    return nil
}

func checkNonZeroVector(v []int) error {
    for _, val := range v {
        if val != 0 {
            return nil
        }
    }
    return fmt.Errorf("vector is a zero vector")
}

func checkMatrixMultiplicationDims(a, b [][]int) error {
    if len(a[0]) != len(b) {
        return fmt.Errorf("matrices have incompatible dimensions for multiplication: %dx%d and %dx%d",
            len(a), len(a[0]), len(b), len(b[0]))
    }
    return nil
}

//func checkMatrixInvertible(m [][]int) error {
//    if err := checkMatrixSquare(m); err != nil {
//        return err
//    }
//    det := calculateDeterminant(m) // Assume this function exists
//    if det == 0 {
//        return fmt.Errorf("matrix is singular and cannot be inverted")
//    }
//    return nil
//}

func checkIsRowVector(v [][]int) error {
    if len(v) != 1 {
        return fmt.Errorf("input is not a row vector")
    }
    return nil
}

func checkIsColumnVector(v [][]int) error {
    if len(v[0]) != 1 {
        return fmt.Errorf("input is not a column vector")
    }
    return nil
}

func checkIdentityMatrix(m [][]int) error {
    if err := checkMatrixSquare(m); err != nil {
        return err
    }
    for i := range m {
        for j := range m[i] {
            if (i == j && m[i][j] != 1) || (i != j && m[i][j] != 0) {
                return fmt.Errorf("matrix is not an identity matrix")
            }
        }
    }
    return nil
}

func getMatrixDimensions(m [][]int) (int, int) {
    return len(m), len(m[0])
}

