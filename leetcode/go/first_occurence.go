package main

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}

	for i, _ := range haystack {
		if haystack[i] == needle[0] {
			var text string
			for j := i; j <= len(haystack)-1; j++ {
				text += string(haystack[j])
				if text == needle {
					return i
				}
			}
		}
	}

	if string(haystack[len(haystack) - 1]) == needle {
		return len(haystack) - 1
	}

	return -1
}

