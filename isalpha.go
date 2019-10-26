package piscine

//awdadar
func IsAlpha(str string) bool {

	i := []rune(str)
	ans := true
	len := 0

	for range str {
		len++
	}
	for g := 0; g < len; g++ { // мы вызываем из массива отдельный элемент g
		//i[g] g- это отдельный элеиент в массиве рун i
		if (i[g] >= 'a' && i[g] <= 'z') || (i[g] >= 'A' && i[g] <= 'Z') || (i[g] >= '0' && i[g] <= '9') {
			//continue // создет контрольно пропуской пункт который чекает всю строчку
		} else { // и если не соответсвтует
			return false // возвращает ответ false
		}
	}
	return ans //
}
