package matType

type matType struct {
	row      int
	col      int
	tamMat   int
	matFloat []float64
}

func (matrix *matType) setParamsMat(matIn [][]string) {
	matrix.row = len(matIn)
	matrix.col = len(matIn[0])
	matrix.tamMat = matrix.col * matrix.row

}
