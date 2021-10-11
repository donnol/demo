package tool

// incompatible type: cannot use i (variable of type int) as K value
// func Map[M map[K]V, K ~int, V string|int](s []V) M {
func Map[M map[int]V, V string|int](s []V) M {
	var m = make(M)
	for i, one := range s {
		m[i] = one
	}
	return m
}
