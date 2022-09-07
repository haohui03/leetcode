package array

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}
	lineMat := len(mat[0])
	output := make([][]int, r)
	for i := range output {
		output[i] = make([]int, c)
	}
	for row := 0; row < r; row++ {
		for line := 0; line < c; line++ {
			i := (row*c + line) / lineMat
			j := (row*c + line) % lineMat
			output[row][line] = mat[i][j]
		}
	}
	return output
}
