package piscine

func BasicAtoi(s string) int {
	var count, sum int        // создаем 2 переменных для того чтобы
	for _, value := range s { // разбили стринг s на руны потому что стринг это уже масив рун
		if value >= '0' && value <= '9' { // делаем рамки работы рун от нуля до девяти
			count = 0                      // чтобы отсчет начинался по новой
			for i := '0'; i < value; i++ { //через фор превращаем руны в интеджер посредством каунтера
				count++
			}
			sum = sum*10 + count //
		}
	}
	return sum
}
