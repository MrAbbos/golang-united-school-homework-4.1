package reverse_string

func ReverseString(input string) (output string) {
	var out = []byte{}
	for i := len(input) - 1; i >= 0; i-- {
		out = append(out, input[i])
	}
	output=string(out)
	return output
}
