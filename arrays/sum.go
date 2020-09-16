package arrays

func Sum(arr []int) (sum int) {
	for _, number := range arr {
		sum += number

	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if (len(numbers) == 0) {
			sums = append(sums, 0)
		} else {

		}

	}
	return sums
}
