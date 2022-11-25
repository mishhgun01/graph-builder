package graph

import "strconv"

func AutoFromAdjMatrix(matrix AdjacencyMatrix) Abstract {
	var res Weighted
	var weighted bool
	for i := 0; i < len(matrix.matrix[0]); i++ {
		for j := i + 1; j < len(matrix.matrix[0]); j++ {
			if matrix.matrix[i][j] != 0 {
				if matrix.matrix[i][j] != 1 {
					weighted = true
				}
				res.Graph[strconv.Itoa(i)][strconv.Itoa(j)] = matrix.matrix[i][j]
			}
		}
	}
	if !weighted {
		return ConvWeightedToUnweighted(&res)
	}
	return res
}
