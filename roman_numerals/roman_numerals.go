package romannumerals

func multiply_letter(number int, letter rune) string {
	str := ""
	for i := 0; i < number; i++ {
		str = str + string(letter)
	}

	return str
}

func ConvertToRoman(arabic int) string {
	str := ""
	if arabic < 4 {
		str += multiply_letter(arabic, 'I')
	}

	return str
}
