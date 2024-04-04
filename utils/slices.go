package utils

// Map applies the given function to each element of the given slice and returns a new slice with the results
func Map[T, U any](arr []T, f func(T) U) []U {
	mapped := make([]U, len(arr))
	for i, elem := range arr {
		mapped[i] = f(elem)
	}
	return mapped
}
