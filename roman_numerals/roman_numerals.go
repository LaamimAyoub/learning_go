package romannumerals

var mapNumbers = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "DM",
	1000: "M",
}

var numbers = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func ConvertToRoman(arabic int) string {
	str := ""
	var temp int
	var res int = arabic

	for _, i := range numbers {

		if res >= i {
			temp = res / i

			res = res - temp*i

			for j := 0; j < temp; j++ {
				str += mapNumbers[i]
			}

		}
	}

	return str
}
