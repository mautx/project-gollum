package matGollum

import (
	"encoding/csv"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
	"os"
	"strconv"
)

// ToMatrix recibe un string con la ruta
// de un archivo CSV para convertirlo en
// una DenseMatrix de la librería gonum
func ToMatrix(str string) *mat.Dense {
	file, err := os.Open(str)
	if err != nil {
		log.Fatal("No es posible encontrar archivo: "+str, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("No se pudo cerrar, wtf, arregla tu vida: "+str, err)
		}
	}(file)

	csvRead := csv.NewReader(file)
	newMatString, err := csvRead.ReadAll()

	if err != nil {
		log.Fatal("No es posible convertir este archivo a matríz"+str, err)
	}
	row := len(newMatString)
	col := len(newMatString[0])
	tamMat := row * col
	newMatFloat := make([]float64, 0, tamMat)

	for _, sliceS := range newMatString {
		for _, element := range sliceS {
			j, err := strconv.ParseFloat(element, 64)

			if err != nil {
				panic(err)
			}
			newMatFloat = append(newMatFloat, j)
		}
	}
	newMatMat := mat.NewDense(row, col, newMatFloat)

	fmt.Println("Matríz parseada: ")
	fmt.Println(mat.Formatted(newMatMat))
	return newMatMat
}
