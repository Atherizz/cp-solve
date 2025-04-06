package main

func MovingShift(s string, shift int) string {
	res := ""

	for _, char := range s {
		value := char + rune(shift)
		if char >= 'a' && char <= 'z' {
			if value > 'z' {
				value = 'a' + (value - 'z' - 1)
			}
			res += string(value)
		} else if char >= 'A' && char <= 'Z' {
			if value > 'Z' {
				value = 'A' + (value - 'Z' - 1)
			}
			res += string(value)
		} else {
			res += string(char)
		}

		shift++
	}

	return res
}

// func DemovingShift(arr []string, shift int) string {
// 	// your code
// }
