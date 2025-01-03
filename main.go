package main

import (
	"fmt"
    "nullspace/nullspace"
)


func main(){ 
	fmt.Println("Result:")
    nullspace.PrintMatrix(nullspace.CreateTriangular(10))
}


