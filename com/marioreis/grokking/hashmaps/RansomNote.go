package hashmaps

func CanConstruct(ransomNote string, magazine string) bool {
	mapMagazine := make(map[rune]int)

	for _, c := range magazine {
		mapMagazine[rune(c)]++
	}

	for _, c := range mapMagazine {
		mapMagazine[rune(c)]--

		if mapMagazine[rune(c)] < 0 {
			return false
		}
	}

	return true
}
