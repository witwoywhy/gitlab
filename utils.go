package gitlab

import "unicode/utf8"

func ToPointer[T any](v T) *T {
	return &v
}

func IsEmptyString(s string) bool {
	return utf8.RuneCountInString(s) == 0
}