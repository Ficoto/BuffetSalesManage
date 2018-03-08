package utils

func IntAbs(a int) int {
	return (a ^ a>>31) - a>>31
}

func Int64Abs(a int64) int64 {
	return (a ^ a>>63) - a>>63
}
