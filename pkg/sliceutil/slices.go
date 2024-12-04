package sliceutil

// Reverse reverses a slice
func Reverse[T any](s []T) []T {
	reversed := make([]T, len(s))
	copy(reversed, s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return reversed
}

// Equal compares two slices contents and returns true if they are the same, false if they are not
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CountDuplicates[T comparable](s []T) int {
	duplicates := 0
	counted := Counts(s)
	for c := range counted {
		if counted[c] > 1 {
			duplicates++
		}
	}
	return duplicates
}

func Counts[V comparable](vs []V) map[V]int {
	return CountsFunc(vs, func(v V) V { return v })
}

func CountsFunc[V any, K comparable](vs []V, key func(V) K) map[K]int {
	h := make(map[K]int)
	for _, v := range vs {
		h[key(v)]++
	}
	return h
}
