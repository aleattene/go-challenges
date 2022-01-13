package solution_sum_squared_numbers

func SumSquaredNumbers(numbers []int) int {
	count := 0
	for i := 0; i < len(numbers); i++ {
		count += numbers[i] * numbers[i]
	}
	return count
}
