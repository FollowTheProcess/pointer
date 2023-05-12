// Package pointer implements a number of helpful and simple tools for working safely and conveniently with pointers.
package pointer

// Or returns the value p is pointing to if it is not nil, else v.
//
//	var s *string // nil pointer
//	str := pointer.Or(s, "default")
//	fmt.Println(str) // "default"
func Or[T any](p *T, v T) T {
	if p != nil {
		return *p
	}
	return v
}
