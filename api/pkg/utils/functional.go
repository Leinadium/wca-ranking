package utils

func Map[T, U any](in []T, f func(T) U) []U {
	var r []U
	for _, v := range in {
		r = append(r, f(v))
	}
	return r
}

func MapNotNull[T, U any](in []*T, f func(*T) U) []U {
	var r []U
	for _, v := range in {
		if v != nil {
			r = append(r, f(v))
		}
	}
	return r
}

func Filter[T any](in []T, f func(T) bool) []T {
	var r []T
	for _, v := range in {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func FilterMap[K comparable, V any](in map[K]V, f func(K, V) bool) []K {
	var r []K

	if in == nil {
		return r
	}

	for k, v := range in {
		if f(k, v) {
			r = append(r, k)
		}
	}
	return r
}

func Flatten[T, U any](in []T, f func(T) []U) []U {
	var r []U
	for _, value := range in {
		r = append(r, f(value)...)
	}
	return r
}

func Last[E any](s []E) (E, bool) {
	if len(s) == 0 {
		var zero E
		return zero, false
	}
	return s[len(s)-1], true
}
