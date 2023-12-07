package sortedmap

// insertInterface inserts the interface{} value v into slice s, at index i.
// and then returns an updated reference.
func insertInterface[V any](s []V, v V, i int) []V {
	s = append(s, *new(V))
	copy(s[i+1:], s[i:])
	s[i] = v

	return s
}

// deleteInterface deletes an interface{} value from slice s, at index i,
// and then returns an updated reference.
func deleteInterface[V any](s []V, i int) []V {
	copy(s[i:], s[i+1:])
	s[len(s)-1] = *new(V)

	return s[:len(s)-1]
}
