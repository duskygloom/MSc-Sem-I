package hashing

func countDigits(n int) int {
	d := 0
	for ; n > 0; n /= 10 {
		d++
	}
	return d
}

func pow10(n int) int {
	r := 1
	for ; n > 0; n-- {
		r *= 10
	}
	return r
}

func MidSquareMethod(key, limit int) int {
	sq := key * key
	sqCount := countDigits(sq)
	limitCount := countDigits(limit)
	if sqCount <= limitCount {
		return sq % limit
	}
	trailing := (sqCount - limitCount) / 2
	sq = sq % pow10(trailing)
	return sq % pow10(limitCount) % limit
}
