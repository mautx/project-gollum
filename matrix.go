package matrix

import (
	"encoding/csv"
	"gollum/matType"
	"gonum.org/v1/gonum/mat"
	"log"
	"os"
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
		log.Fatal("No es posible convertir este archivo")
	}

	newMatMat := matType.MatType{}
	newMatMat.NewMat(newMatString)

	return newMatMat.GetMat()

}
