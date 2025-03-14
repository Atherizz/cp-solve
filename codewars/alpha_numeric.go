package main

func Alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}	

	for _, char := range str {
		if !(char >= '0' && char <= '9') && !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z') {
			return false
		}
	}

return true
}
