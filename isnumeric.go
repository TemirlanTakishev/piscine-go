package piscine

func IsNumeric(str string) bool {

	Per := []rune(str) // переобразуем в руну
	len := 0
	ans := true

	for range str { // формула чтобы комп видел строку
		len++
	}
	for g := 0; g < len; g++ { // делим массив на отдельные элементы
		if Per[g] >= '0' && Per[g] <= '9' { // говорим что прога будет работать в заданных значениях
			continue // фейс контроль
		} else {
			ans = false
		}

	}
	return ans
}
