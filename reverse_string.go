package reverse_string

func ReverseString(input string) (output string) {
	var out = []rune{}
	var out2 = []rune{}
	for _, a := range input {
		out = append(out, a)
	}

	for i := len(out) - 1; i >= 0; i-- {
		out2 = append(out2, out[i])
	}

	output = string(out2)
	return output
}
