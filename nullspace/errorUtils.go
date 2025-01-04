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

