package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {

	rows := strings.Split(s, "\n")
	if strings.Count(s, "\n") != len(rows)-1 {
		return nil, fmt.Errorf("Row cannot be empty")
	}
	matrix := make(Matrix, len(rows))
	nCols := 0
	for i, row := range rows {
		cols := strings.Split(strings.Trim(row, " "), " ")
		if nCols != len(cols) {
			if nCols == 0 {
				nCols = len(cols)
			} else {
				return nil, fmt.Errorf("Uneven rows")
			}
		}
		line := make([]int, nCols)
		for j, col := range cols {
			var err error
			if line[j], err = strconv.Atoi(col); err != nil {
				return nil, err
			}
		}
		matrix[i] = line
	}
	return matrix, nil
}

func (m Matrix) Rows() [][]int {
	matrix := make([][]int, len(m))
	for i, row := range m {
		matrix[i] = make([]int, len(row))
		copy(matrix[i], m[i])
	}
	return matrix
}

func (m Matrix) Cols() [][]int {
	matrix := make([][]int, len(m[0]))
	for i := 0; i < len(m[0]); i++ {
		matrix[i] = make([]int, len(m))
		for j := 0; j < len(m); j++ {
			matrix[i][j] = m[j][i]
		}
	}
	return matrix
}

func (m Matrix) Set(row, col int, val int) bool {
	if !(row >= 0 && row < len(m)) || !(col >= 0 && col < len(m[0])) {
		return false
	}
	m[row][col] = val
	return true
}
