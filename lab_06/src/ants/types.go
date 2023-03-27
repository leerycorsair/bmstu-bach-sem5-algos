package ants

type MatrixInt [][]int

type ParamsType struct {
	alpha   float64
	beta    float64
	q       float64
	p       float64
	ph_coef float64
	days    int
}

type ColonyType struct {
	mtr      MatrixInt
	pheromon [][]float64
	params   ParamsType
}

type AntType struct {
	colony  *ColonyType
	visited MatrixInt
	path    []int
	vertex  int
}
