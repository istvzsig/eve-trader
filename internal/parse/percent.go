package parse

import "strconv"

func PercentArg(s string) (float64, error) {
	// Accept: "15", "15%", "20.5", "20.5%"
	if len(s) > 0 && s[len(s)-1] == '%' {
		s = s[:len(s)-1]
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}
