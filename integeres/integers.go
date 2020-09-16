package integers

func Add(a, b int) (sum int) {
	sum = a + b
	return
}

func Repeat(str string, times int) (result string) {
	for i := 0; i < times; i++ {
		result += str
	}

	return
}
