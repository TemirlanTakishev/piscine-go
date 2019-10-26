package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {

		z01.PrintRune('0')
	}
	array := []int{}

	for n != 0 { // н не равно нулю потому что когда то она не не равно нулю
		ostatok := n % 10              // создали для того чтобы сохранить в array
		array = append(array, ostatok) // сохраняем в масив array

		// убрать последнию цифру з числа  n
		n = n / 10
	}

	// fmt.Println(array)

	len := 0          //  создаем его для того чтобы вычислить длину
	for range array { // проходиться по всей длинне масива
		len++ // прибавить одн
	}

	// sortirovka
	for k := 1; k < len; k++ { // создаем цикл который должен пройти от начало до конца len раз
		for i := 1; i < len; i++ { // еще раз создаем цикл пройти от начало до конца len раз (если мы пройдем один раз то поставит только самое большое число назад но другие остануться в своем порядке  )
			if array[i-1] > array[i] { // проверка если левое больше правого то
				array[i-1], array[i] = array[i], array[i-1] // меняем левое с правым
			}
		}
	}

	// vivod
	for _, nb := range array { // проходиться по всему масиву
		z01.PrintRune(rune(nb + 48)) // принтит (оборачиваем в руну это называется каст)
	}

	// fmt.Println(array)
}

//Write a function which prints the digits of an int passed in parameter in ascending order. All possible values of type int have to go through, besides negative numbers. Convertion to int64 is not allowed.
