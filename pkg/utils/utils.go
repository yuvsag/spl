package utils

import "golang.org/x/exp/constraints"

func SliceString(s string, n int) []string {
	result := []string{}
	for i := 0; i < len(s); i += n {
		end := Min(i+n, len(s))

		result = append(result, s[i:end])
	}
	return result
}
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
