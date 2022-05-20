package matrix_utils

func IsSquare(m []string) bool {
	for i := 0; i < len(m); i++ {
		if len(m[i]) != len(m) {
			return false
		}
	}
	return true
}
