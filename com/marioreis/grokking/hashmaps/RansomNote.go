package hashmaps

func CanConstruct(ransomNote string, magazine string) bool {
	mapMagazine := make(map[rune]int)
	mapRansomNote := make(map[rune]int)

	for _, c := range magazine {
		mapMagazine[rune(c)]++
	}

	for _, c := range ransomNote {
		mapRansomNote[rune(c)]++
	}

	for k, _ := range mapRansomNote {
		if mapRansomNote[k] > mapMagazine[k] {
			return false
		}
	}

	return true
}
