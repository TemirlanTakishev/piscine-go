package piscine

func Capitalize(s string) string {

	sring := []rune(s) // создаем массив рун cast

	len := 0 // вычисляем длину строки
	for range sring {
		len++
	}
	for i, bykva := range sring { // даёт доступ к каждой букве и диджителу

		if i == 0 || !isAlphaNum(sring[i-1]) { // проверяем если первая буква в слове

			if bykva >= 'a' && bykva <= 'z' { // проверяем если маленькая
				sring[i] = bykva - 32 //то переварачиваем в большую
			}
		} else {
			if bykva >= 'A' && bykva <= 'Z' { // если это большая делаем ее маленькую
				sring[i] = bykva + 32 //делаем ее маленькую
			}
		}

	}

	return string(sring) // возвращаем стринг

}

func isAlphaNum(a rune) bool { // сделали чтобы проверить

	if a >= 'a' && a <= 'z' { // маленькая ли буква
		return true
	}
	if a >= 'A' && a <= 'Z' { // большая ли буква
		return true
	}
	if a >= '0' && a <= '9' { // цифра ли
		return true
	}

	return false
}
