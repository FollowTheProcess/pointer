// Package pointer is a tiny, simple and obvious library to help deal safely with pointers
package pointer

// New returns a pointer to the passed value, particularly useful for
// generating pointers to primitive types (string, int etc.).
//
//	s := pointer.New("hello")
//	fmt.Printf("%T\n", s) // *string
func New[T any](thing T) *T {
	return &thing
}

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

// OrDefault returns the value p is pointing to if it is not nil, otherwise
// the zero value for the type T.
//
//	var s *string // nil pointer
//	str := pointer.OrDefault(s)
//	fmt.Println(str) // ""
func OrDefault[T any](p *T) T {
	return Or(p, *new(T))
}
