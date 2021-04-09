package main 

func Sum(numbers[3]int) int {
	sum := 0
	for _, number := range numbers{
		sum  += number
	}
	return sum
}