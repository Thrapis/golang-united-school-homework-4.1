package reverse_string

func ReverseString(input string) (output string) {
	var buf = []rune(input)
	for i := 0; i < len(buf)-1-i; i++ {
		buf[i], buf[len(buf)-1-i] = buf[len(buf)-1-i], buf[i]
	}
	output = string(buf)
	return output
}
