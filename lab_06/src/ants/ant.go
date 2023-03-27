package ants

import (
	"math"
	"math/rand"
	"time"
)

const (
	DEFAULT_ALPHA        = 3.0
	DEFAULT_BETA         = 7.0
	DEFAULT_P            = 0.6
	DEFAULT_PHEROMONCOEF = 0.5
	DEFAULT_DAYS         = 5
)

func AntAlgo(citiesMap MatrixInt, alpha float64, beta float64, p float64, ph_coef float64, days int) (int, []int) {
	params := new(ParamsType)
	params.alpha = alpha
	params.beta = beta
	params.p = p
	params.ph_coef = ph_coef
	params.days = days
	params.q = float64(calcQ(citiesMap))
	colony := createColony(citiesMap, params)
	return live(colony)
}

func calcQ(citiesMap MatrixInt) float64 {
	sum := float64(0)
	for i := 0; i < len(citiesMap); i++ {
		for j := 0; j < len(citiesMap); j++ {
			sum += float64(citiesMap[i][j])
		}
	}
	sum /= float64(len(citiesMap))
	return sum
}

func createColony(citiesMap MatrixInt, params *ParamsType) *ColonyType {
	n := len(citiesMap)
	colony := new(ColonyType)
	colony.mtr = citiesMap
	colony.params = *params
	colony.pheromon = make([][]float64, n)
	for i := 0; i < n; i++ {
		colony.pheromon[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			colony.pheromon[i][j] = colony.params.ph_coef
		}
	}
	return colony
}

func live(colony *ColonyType) (int, []int) {
	n := len(colony.mtr)
	min_dist := 0
	min_path := make([]int, n)
	for i := 0; i < colony.params.days; i++ {
		for j := 0; j < n; j++ {
			ant := colony.createAnt(j)
			ant.findPath()
			ant.updatePheromon(colony.params)
			cur_dist := ant.distance()
			if (cur_dist < min_dist) || (min_dist == 0) {
				min_dist = cur_dist
				copy(min_path, ant.path)
			}
		}
	}
	return min_dist, min_path
}

func (ant *AntType) findPath() {
	for {
		p := ant.getProbabilities()
		vertex := chooseVertex(p)
		if vertex != -1 {
			ant.move(vertex)
		} else {
			break
		}
	}
}

func (colony *ColonyType) createAnt(vertex int) *AntType {
	n := len(colony.mtr)
	ant := new(AntType)
	ant.colony = colony
	ant.visited = make(MatrixInt, n)
	for i := 0; i < n; i++ {
		ant.visited[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ant.visited[i][j] = colony.mtr[i][j]
		}
	}
	ant.vertex = vertex
	ant.path = make([]int, 0)
	ant.path = append(ant.path, ant.vertex)
	return ant
}

func (ant *AntType) getProbabilities() []float64 {
	n := len(ant.colony.mtr)
	probs := make([]float64, n)
	psum := float64(0)
	var (
		param1 float64
		param2 float64
	)
	for i := 0; i < n; i++ {
		if ant.visited[ant.vertex][i] != 0 {
			param1 = ant.colony.pheromon[ant.vertex][i]
			param2 = (float64(1) /
				float64(ant.colony.mtr[ant.vertex][i]))
			d := math.Pow(param1, ant.colony.params.beta) *
				math.Pow(param2, ant.colony.params.alpha)
			probs[i] = d
			psum += d
		} else {
			probs[i] = 0
		}
	}
	for i := 0; i < n; i++ {
		probs[i] /= psum
	}
	return probs
}

func chooseVertex(p []float64) int {
	n := len(p)
	sum := float64(0)
	for i := 0; i < n; i++ {
		sum += p[i]

	}
	randomProbability :=
		rand.New(rand.NewSource(time.Now().UnixNano())).Float64() * sum
	sum = 0
	for i := 0; i < n; i++ {
		if randomProbability < sum+p[i] && randomProbability > sum {
			return i
		}
		sum += p[i]
	}
	return -1
}

func (ant *AntType) move(new_vertex int) {
	n := len(ant.colony.mtr)
	for i := 0; i < n; i++ {
		ant.visited[i][ant.vertex] = 0
	}
	ant.path = append(ant.path, new_vertex)
	ant.vertex = new_vertex
}

func (ant *AntType) distance() int {
	distance := 0
	n := len(ant.path)
	for i := 0; i < n; i++ {
		src := ant.path[i]
		dst := ant.path[(i+1)%len(ant.path)]
		distance += ant.colony.mtr[src][dst]
	}
	return distance
}

func (ant *AntType) updatePheromon(params ParamsType) {
	delta := float64(0)
	for i := 0; i < len(ant.colony.pheromon); i++ {
		for j := 0; j < len(ant.colony.pheromon); j++ {
			if ant.colony.mtr[i][j] != 0 {
				f := false
				for k := 0; k < len(ant.path); k++ {
					src := ant.path[k]
					dst := ant.path[(k+1)%len(ant.path)]
					if (src == i && dst == j) || (dst == i && src == j) {
						f = true
					}
				}
				if f {
					delta += ant.colony.params.q /
						float64(ant.colony.mtr[i][j])
				}

			}
			ant.colony.pheromon[i][j] = (1 - params.p) *
				(float64(ant.colony.pheromon[i][j]) + delta)
			if ant.colony.pheromon[i][j] <= 0 {
				ant.colony.pheromon[i][j] = 0.1
			}
		}

	}
}


