package piscine

func ConcatParams(args []string) string {
	result := ""
	k := 0

	for index := range args {
		k = index
	}

	for index, value := range args {
		result += value
		if index != k {
			result += "\n"
		}
	}

	return result
}
