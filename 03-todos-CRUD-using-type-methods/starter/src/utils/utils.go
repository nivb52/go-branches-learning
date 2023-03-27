package utils

func SlicesDeleteFast[T any](s []T, index int) []T {
	if index >= len(s) || index < 0 {
		return s
	}
	if index == len(s)-1 {
		return s[:len(s)-1]
	}
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}

// func generateUUID () string {
// 	return time.Now().UnixNano()
// }
