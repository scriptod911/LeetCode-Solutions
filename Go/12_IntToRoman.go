func intToRoman(num int) string {
	var roman string
	var n int
	for i := 0; i < 4; i++ {
		n = num % 10
		num /= 10
		switch i {
		case 0:
			roman = romanIntToRoman(n, "I", "V", "X") + roman
		case 1:
			roman = romanIntToRoman(n, "X", "L", "C") + roman
		case 2:
			roman = romanIntToRoman(n, "C", "D", "M") + roman
		case 3:
			roman = romanIntToRoman(n, "M", "", "") + roman
		}
	}
	return roman
}


func romanIntToRoman(n int, one string, five string, ten string) string {
	var roman string
	switch n {
	case 0:
		roman = ""
	case 1:
		roman = one
	case 2:
		roman = one + one
	case 3:
		roman = one + one + one
	case 4:
		roman = one + five
	case 5:
		roman = five
	case 6:
		roman = five + one
	case 7:
		roman = five + one + one
	case 8:
		roman = five + one + one + one
	case 9:
		roman = one + ten
	}
	return roman
}