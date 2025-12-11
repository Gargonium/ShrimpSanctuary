package utils

func ClampAndBounce(pos, min, max, vel int32) (int32, int32) {
	if pos < min {
		return min, -vel
	}
	if pos > max {
		return max, -vel
	}
	return pos, vel
}
