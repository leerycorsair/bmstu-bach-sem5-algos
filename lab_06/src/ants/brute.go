package ants

func BruteForceAlgo(citiesMap MatrixInt, cities int) (int, []int) {
	path := make([]int, cities)
	for i := 0; i < cities; i++ {
		path[i] = i
	}

	minLen := dist(citiesMap, path)
	minPath := make([]int, cities)

	currPath, state := path, true

	for state {
		currLen := dist(citiesMap, currPath)
		if currLen < minLen {
			minLen = currLen
			copy(minPath, currPath)
		}
		currPath, state = haveNextPermutation(currPath)
	}

	return minLen, minPath
}

func haveNextPermutation(path []int) ([]int, bool) {
	j := len(path) - 2
	for j != -1 && path[j] >= path[j+1] {
		j--
	}
	if j == -1 {
		return path, false
	}

	k := len(path) - 1
	for path[j] >= path[k] {
		k--
	}
	swap(path, j, k)
	l := j + 1
	r := len(path) - 1

	for l < r {
		swap(path, l, r)
		l++
		r--
	}
	return path, true
}

func swap(path []int, i int, j int) {
	s := path[i]
	path[i] = path[j]
	path[j] = s
}

func dist(citiesMap MatrixInt, path []int) int {
	dist := 0
	for k := 0; k < (len(path) - 1); k++ {
		dist += citiesMap[path[k]][path[k+1]]
	}
	dist += citiesMap[path[0]][path[len(path)-1]]
	return dist
}
