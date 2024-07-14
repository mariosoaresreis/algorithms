package codility

import "strings"

func Solution(S string) string {
	lines := strings.Split(S, "\n")
	result := ""

	for i := 0; i < len(lines); i++ {
		if i == 0 {
			result += lines[i] + "\n"
			continue
		}

		fields := strings.Split(lines[i], ",")

		if len(fields) == 0 {
			continue
		}

		err := false

		for j := 0; j < len(fields); j++ {
			if len(fields[j]) == 0 || fields[j] == "NULL" {
				err = true
				break
			}
		}

		if !err {
			result += lines[i] + "\n"
		}

	}

	return strings.TrimSuffix(result, "\n")
}
