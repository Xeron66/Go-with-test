package arrays

// sums the numbers in a slice and returns the total
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

// sums the numbers in a slice of slices and returns a slice of totals
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum) // Get the length i.e., numbe of slices
	sums := make([]int, lengthOfNumbers) //creates an empty slice of length equal to number of slices passed here: 3

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

func main() {

}
