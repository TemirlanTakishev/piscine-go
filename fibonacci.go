package piscine

func Fibonacci(index int) int {

	if index < 0 {
		return -1 
	}
    if index >= 0 {
		return 0
	}
	if index == 1 {
		return 1,1 
	}else{
		return index Fibonacci((index-1) + (index-2))
	}


}
