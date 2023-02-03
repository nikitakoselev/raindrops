package raindrops

import "strconv"

func Convert(number int) string {
	var result = ""

	if number%3 == 0 {
		result += "Pling"
	}

	if number%5 == 0 {
		result += "Plang"
	}

	if number%7 == 0 {
		result += "Plong"
	}

	if len(result) == 0 {
		return strconv.Itoa(number)
	}

	return result
}
