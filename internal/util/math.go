package util

func MinValue(values ...int) int {

	result := values[0]

	for _, v := range values[1:] {

		if v < result {
			result = v
		}
	}

	return result
}
