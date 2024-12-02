package utils

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SignInt(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}
