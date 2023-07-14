package utils

func MapSlice[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for idx, val := range a {
		n[idx] = f(val)
	}
	return n
}
