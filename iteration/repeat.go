package iteration

// Repeats the character 5 times
func Repeat(character string, num int) (repeated string) {
	for i := 0; i < num; i++ {
		repeated += character
	}
	return
}
