package integers

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}

func Repeat(character string, times int) string {
	repeated := ""
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
