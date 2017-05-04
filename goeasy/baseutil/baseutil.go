package baseutil

func Abs(value int) int {
	if value >= 0 {
		return value
	} else {
		return -value
	}
}
