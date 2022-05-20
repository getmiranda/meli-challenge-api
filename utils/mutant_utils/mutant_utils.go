package mutant_utils

import "fmt"

var (
	mutantDna = []string{"AAAA", "TTTT", "CCCC", "GGGG"}
)

type Coordinate struct {
	X int
	Y int
}

type Result struct {
	IsMutant    bool
	Dna         string
	Coordinates []Coordinate
}

func IsSquare(m []string) bool {
	for i := 0; i < len(m); i++ {
		if len(m[i]) != len(m) {
			return false
		}
	}
	return true
}

func IsMutant(dna []string) *Result {
	for i := 0; i < len(dna); i++ {
		result := isHorizontal(dna, i)
		if result != nil && result.IsMutant {
			return result
		}

		result = isVertical(dna, i)
		if result != nil && result.IsMutant {
			return result
		}

		result = isDiagonal(dna, i)
		if result != nil && result.IsMutant {
			return result
		}
	}
	return &Result{IsMutant: false}
}

func isHorizontal(dna []string, i int) *Result {
	for j := 0; j < len(dna[i]); j++ {
		if j+3 < len(dna[i]) {
			if dna[i][j] == dna[i][j+1] && dna[i][j] == dna[i][j+2] && dna[i][j] == dna[i][j+3] {
				dnaString := dnaString(dna[i][j], dna[i][j+1], dna[i][j+2], dna[i][j+3])
				coord := []Coordinate{{i, j}, {i, j + 1}, {i, j + 2}, {i, j + 3}}
				if !isMutantDna(dnaString) {
					continue
				}
				return &Result{true, dnaString, coord}
			}
		}
	}
	return nil
}

func isVertical(dna []string, i int) *Result {
	for j := 0; j < len(dna); j++ {
		if j+3 < len(dna) {
			if dna[j][i] == dna[j+1][i] && dna[j][i] == dna[j+2][i] && dna[j][i] == dna[j+3][i] {
				dnaString := dnaString(dna[j][i], dna[j+1][i], dna[j+2][i], dna[j+3][i])
				coord := []Coordinate{{j, i}, {j + 1, i}, {j + 2, i}, {j + 3, i}}
				if !isMutantDna(dnaString) {
					continue
				}
				return &Result{true, dnaString, coord}
			}
		}
	}
	return nil
}

func isDiagonal(dna []string, i int) *Result {
	for j := 0; j < len(dna); j++ {
		if j+3 < len(dna) && i+3 < len(dna) {
			if dna[j][i] == dna[j+1][i+1] && dna[j][i] == dna[j+2][i+2] && dna[j][i] == dna[j+3][i+3] {
				dnaString := dnaString(dna[j][i], dna[j+1][i+1], dna[j+2][i+2], dna[j+3][i+3])
				coord := []Coordinate{{j, i}, {j + 1, i + 1}, {j + 2, i + 2}, {j + 3, i + 3}}
				if !isMutantDna(dnaString) {
					continue
				}
				return &Result{true, dnaString, coord}
			}
		}
	}
	for j := 0; j < len(dna); j++ {
		if j+3 < len(dna) && i-3 >= 0 {
			if dna[j][i] == dna[j+1][i-1] && dna[j][i] == dna[j+2][i-2] && dna[j][i] == dna[j+3][i-3] {
				dnaString := dnaString(dna[j][i], dna[j+1][i-1], dna[j+2][i-2], dna[j+3][i-3])
				coord := []Coordinate{{j, i}, {j + 1, i - 1}, {j + 2, i - 2}, {j + 3, i - 3}}
				if !isMutantDna(dnaString) {
					continue
				}
				return &Result{true, dnaString, coord}
			}
		}
	}
	return nil
}

func isMutantDna(dns string) bool {
	for _, dna := range mutantDna {
		if dns == dna {
			return true
		}
	}
	return false
}

func dnaString(s1, s2, s3, s4 byte) string {
	return fmt.Sprintf("%s%s%s%s", string(s1), string(s2), string(s3), string(s4))
}
