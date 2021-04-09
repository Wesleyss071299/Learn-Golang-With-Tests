package main 

func Sum(numbers[3]int) int {
	sum := 0
	for i:=0; i < 3; i++ {
		sum  += numbers[i]
	}
	return sum
}