package sortingalgorithms

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	return arr
}

func MergeSort(arr []int) []int {

	return sort(arr, 0, len(arr)-1)
}

func sort(arr []int, start, end int) []int {

	if end > start {
		middle := (start + end) / 2

		sort(arr, start, middle)
		sort(arr, middle+1, end)

		merge(arr, start, middle, end)

	}

	return arr
}

func merge(arr []int, start, middle, end int) []int {

	left := arr[start : middle+1]
	right := arr[middle+1 : end+1]

	leftArr := make([]int, len(left))
	rightArr := make([]int, len(right))

	copy(leftArr[:], left)
	copy(rightArr[:], right)

	var i, j int
	k := start

	for i < len(right) && j < len(left) {
		if rightArr[i] < leftArr[j] {
			arr[k] = rightArr[i]
			i++
		} else {
			arr[k] = leftArr[j]
			j++
		}
		k++
	}

	for i < len(right) {
		arr[k] = rightArr[i]
		i++
		k++
	}

	for j < len(left) {
		arr[k] = leftArr[j]
		j++
		k++
	}

	return arr
}
