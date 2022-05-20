package main

// Recibirás como parámetro un array de Strings que representan cada fila de una tabla de (NxN) con la secuencia del ADN. Las letras de los Strings solo pueden ser: (A,T,C,G), las cuales representa cada base nitrogenada del ADN.
//
// Sabrás si un humano es mutante, si encuentras ​más de una secuencia de cuatro letras iguales​, de forma oblicua, horizontal o vertical.
//
// Ejemplo
//
// Entrada:
//
// dna := []string{"ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"}
//
// Salida:
//
// En este caso el llamado a la función isMutant(dna) devuelve “true”.
//
// Desarrolla el algoritmo de la manera más eficiente posible.
//
/*

func isMutant(dna []string) bool {
	if !isSquare(dna) {
		fmt.Println("Not a square")
		return false
	}
	printMatrix(dna)
	for i := 0; i < len(dna); i++ {
		if isHorizontal(dna, i) {
			fmt.Println("isHorizontal")
			return true
		}
		if isVertical(dna, i) {
			fmt.Println("isVertical")
			return true
		}
		if isDiagonal(dna, i) {
			fmt.Println("isDiagonal")
			return true
		}
	}
	return false
}

func isHorizontal(dna []string, i int) bool {
	for j := 0; j < len(dna[i]); j++ {
		if j+3 < len(dna[i]) {
			if dna[i][j] == dna[i][j+1] && dna[i][j] == dna[i][j+2] && dna[i][j] == dna[i][j+3] {
				return true
			}
		}
	}
	return false
}

func isVertical(dna []string, i int) bool {
	for j := 0; j < len(dna); j++ {
		if j+3 < len(dna) {
			if dna[j][i] == dna[j+1][i] && dna[j][i] == dna[j+2][i] && dna[j][i] == dna[j+3][i] {
				return true
			}
		}
	}
	return false
}

func isDiagonal(dna []string, i int) bool {
	for j := 0; j < len(dna); j++ {
		if j+3 < len(dna) && i+3 < len(dna) {
			if dna[j][i] == dna[j+1][i+1] && dna[j][i] == dna[j+2][i+2] && dna[j][i] == dna[j+3][i+3] {
				fmt.Printf("%s(%d,%d) %s(%d,%d) %s(%d,%d) %s(%d,%d)\n", string(dna[j][i]), j, i, string(dna[j+1][i+1]), j+1, i+1, string(dna[j+2][i+2]), j+2, i+2, string(dna[j+3][i+3]), j+3, i+3)
				return true
			}
		}
	}
	for j := 0; j < len(dna); j++ {
		if j+3 < len(dna) && i-3 >= 0 {
			if dna[j][i] == dna[j+1][i-1] && dna[j][i] == dna[j+2][i-2] && dna[j][i] == dna[j+3][i-3] {
				fmt.Printf("%s(%d,%d) %s(%d,%d) %s(%d,%d) %s(%d,%d)\n", string(dna[j][i]), j, i, string(dna[j+1][i-1]), j+1, i-1, string(dna[j+2][i-2]), j+2, i-2, string(dna[j+3][i-3]), j+3, i-3)
				return true
			}
		}
	}
	return false
}

func main() {
	dna := []string{
		"ATGCTA",
		"CTGTAC",
		"TATTGT",
		"ATAAGG",
		"CCGCTA",
		"ATATTC",
	}
	fmt.Println(isMutant(dna))

	dna = []string{
		"ATGCAAG",
		"CTGTACG",
		"TATTGTG",
		"ATAAGGG",
		"CCGCTAG",
		"ATATTCG",
	}
	fmt.Println(isMutant(dna))
}

func printMatrix(dna []string) {
	for i := 0; i < len(dna); i++ {
		for j := 0; j < len(dna[i]); j++ {
			fmt.Printf("%s ", string(dna[i][j]))
		}
		fmt.Println()
	}
}

*/
