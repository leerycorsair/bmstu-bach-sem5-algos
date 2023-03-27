package levenshtein

func LevRec(str1, str2 string) int {
	s1, s2 := []rune(str1), []rune(str2)

	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	return lev_dist_rec(s1, s2, len(s1), len(s2))
}

func lev_dist_rec(s1, s2 []rune, l1, l2 int) int {
	if l1 == 0 {
		return l2
	}

	if l2 == 0 && l1 > 0 {
		return l1
	}

	check := 1
	if s1[l1-1] == s2[l2-1] {
		check = 0
	}

	return min_value(
		lev_dist_rec(s1, s2, l1, l2-1)+1,
		lev_dist_rec(s1, s2, l1-1, l2)+1,
		lev_dist_rec(s1, s2, l1-1, l2-1)+check)
}

func LevRecCached(str1, str2 string) int {
	s1, s2 := []rune(str1), []rune(str2)

	l1, l2 := len(s1), len(s2)
	if l1 == 0 || l2 == 0 {
		return max_value(l1, l2)
	}

	mtr := levmtr_init(l1, l2, true)

	return dist_rec_cached(s1, s2, len(s1), len(s2), mtr)
}

func dist_rec_cached(s1, s2 []rune, l1, l2 int, mtr [][]int) int {
	if l1 == 0 {
		return l2
	}

	if l2 == 0 {
		return l1
	}

	if mtr[l1-1][l2-1] != -1 {
		return mtr[l1-1][l2-1]
	}

	if s1[l1-1] == s2[l2-1] {
		mtr[l1-1][l2-1] = dist_rec_cached(s1, s2, l1-1, l2-1, mtr)
		return mtr[l1-1][l2-1]
	}

	mtr[l1-1][l2-1] = 1 + min_value(
		dist_rec_cached(s1, s2, l1, l2-1, mtr),
		dist_rec_cached(s1, s2, l1-1, l2, mtr),
		dist_rec_cached(s1, s2, l1-1, l2-1, mtr))

	return mtr[l1-1][l2-1]
}

func LevMtr(str1, str2 string) (int, [][]int) {
	s1, s2 := []rune(str1), []rune(str2)
	l1, l2 := len(s1), len(s2)
	mtr := levmtr_init(l1+1, l2+1, false)

	check := 0
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if s1[i-1] == s2[j-1] {
				check = 0
			} else {
				check = 1
			}
			mtr[i][j] = min_value(
				mtr[i-1][j]+1,
				mtr[i][j-1]+1,
				mtr[i-1][j-1]+check)
		}
	}
	dist := mtr[l1][l2]

	return dist, mtr
}

func DamLevMtr(str1, str2 string) (int, [][]int) {
	s1, s2 := []rune(str1), []rune(str2)
	l1, l2 := len(s1), len(s2)
	mtr := levmtr_init(l1+1, l2+1, false)
	check := 0
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if s1[i-1] == s2[j-1] {
				check = 0
			} else {
				check = 1
			}
			mtr[i][j] = min_value(
				mtr[i-1][j]+1,
				mtr[i][j-1]+1,
				mtr[i-1][j-1]+check)

			if i > 1 && j > 1 && s1[i-1] == s2[j-2] && s1[i-2] == s2[j-1] {
				mtr[i][j] = min_value(mtr[i][j], mtr[i-2][j-2]+1)
			}
		}
	}
	dist := mtr[l1][l2]
	return dist, mtr
}

func levmtr_init(n, m int, is_rec bool) [][]int {
	mtr := make([][]int, n)
	for i := range mtr {
		mtr[i] = make([]int, m)
		mtr[i][0] = i
	}

	for j := 1; j < m; j++ {
		mtr[0][j] = j
	}

	if is_rec {
		for i := range mtr {
			for j := range mtr[i] {
				mtr[i][j] = -1
			}
		}
	}

	return mtr
}

func DamLevRec(str1, str2 string) int {
	s1, s2 := []rune(str1), []rune(str2)

	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	return damlev_dist_rec(s1, s2, len(s1), len(s2))
}

func damlev_dist_rec(s1, s2 []rune, l1, l2 int) int {
	if l1 == 0 {
		return l2
	}

	if l2 == 0 && l1 > 0 {
		return l1
	}

	check := 0
	if s1[l1-1] != s2[l2-1] {
		check = 1
	}

	if l1 > 1 && l2 > 1 && s1[l1-1] == s2[l2-2] && s1[l1-2] == s2[l2-1] {
		return min_value(
			damlev_dist_rec(s1, s2, l1-1, l2)+1,
			damlev_dist_rec(s1, s2, l1, l2-1)+1,
			damlev_dist_rec(s1, s2, l1-1, l2-1)+check,
			damlev_dist_rec(s1, s2, l1-2, l2-2)+1)
	} else {
		return min_value(
			damlev_dist_rec(s1, s2, l1-1, l2)+1,
			damlev_dist_rec(s1, s2, l1, l2-1)+1,
			damlev_dist_rec(s1, s2, l1-1, l2-1)+check)

	}
}
