package piscine

func SplitWhiteSpaces(str string) []string {
	k := 0
	bool := false
	for index := range str {
		if bool && index != 0 && (str[index-1] == '\n' || str[index-1] == '\t' || str[index-1] == ' ') && str[index] != '\n' && str[index] != '\t' && str[index] != ' ' {
			k++
		}
		if str[index] != '\n' && str[index] != '\t' && str[index] != ' ' {
			bool = true
		}
	}
	k++

	x := 0
	result := make([]string, k)
	ok := true
	for _, value := range str {
		if value == '\n' || value == '\t' || value == ' ' {
			if !ok {
				x++
			}
			ok = true
			continue
		}
		result[x] = result[x] + string(value)
		ok = false
	}
	return result
}
