package tree

func charToInt(ch byte) int {
	if ch < 'a' || ch > 'z' {
		return 0
	}
	return int(ch - 'a')
}
