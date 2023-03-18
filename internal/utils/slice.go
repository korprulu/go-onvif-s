// Package utils ...
package utils

// Map the map creates a new array populated with the results of calling a
// provided function on every element in the calling array.
func Map[S, M any](s []S, fn func(S) M) []M {
	m := make([]M, len(s))
	for i, element := range s {
		m[i] = fn(element)
	}
	return m
}
