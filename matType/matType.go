package matType

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"strconv"
)

type MatType struct {
	row      int
	col      int
	tamMat   int
	matFloat []float64
	matDense *mat.Dense
}

// NewMat newMat crea una matríz de tipo Dense cuando recibe
// una matríz de tipo String
func (matrix *MatType) NewMat(matIn [][]string) {
	matrix.row = len(matIn)
	matrix.col = len(matIn[0])
	matrix.tamMat = matrix.col * matrix.row
	matrix.matFloat = make([]float64, 0, matrix.tamMat)

	for _, sliceS := range matIn {
		for _, element := range sliceS {
			j, err := strconv.ParseFloat(element, 64)

			if err != nil {
				panic(err)
			}
			matrix.matFloat = append(matrix.matFloat, j)
		}
	}
	matrix.matDense = mat.NewDense(matrix.row, matrix.col, matrix.matFloat)

	fmt.Println("Matríz parseada: ")
	fmt.Println(mat.Formatted(matrix.matDense))

}

// GetMat regresa la matriz tipo Dense
func (matrix MatType) GetMat() *mat.Dense {
	return matrix.matDense
}
