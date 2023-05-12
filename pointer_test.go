package pointer_test

import (
	"fmt"
	"testing"

	"github.com/FollowTheProcess/pointer"
)

func TestOr(t *testing.T) {
	t.Run("string nil", func(t *testing.T) {
		var str *string // nil
		or := "default"

		got := pointer.Or(str, or)

		if got != or {
			t.Errorf("got %q, wanted %q", got, or)
		}
	})

	t.Run("string non nil", func(t *testing.T) {
		str := "hello"
		or := "default"

		got := pointer.Or(&str, or)

		if got != str {
			t.Errorf("got %q, wanted %q", got, str)
		}
	})

	t.Run("int nil", func(t *testing.T) {
		var i *int // nil
		or := 42

		got := pointer.Or(i, or)

		if got != or {
			t.Errorf("got %d, wanted %d", got, or)
		}
	})

	t.Run("int non nil", func(t *testing.T) {
		i := 12
		or := 42

		got := pointer.Or(&i, or)

		if got != i {
			t.Errorf("got %d, wanted %d", got, i)
		}
	})

	t.Run("struct nil", func(t *testing.T) {
		type thing struct{ name string }
		var x *thing
		or := thing{name: "default"}

		got := pointer.Or(x, or)

		if got != or {
			t.Errorf("got %v, wanted %v", got, or)
		}
	})

	t.Run("struct non nil", func(t *testing.T) {
		type thing struct{ name string }
		x := thing{name: "x"}
		or := thing{name: "default"}

		got := pointer.Or(&x, or)

		if got != x {
			t.Errorf("got %v, wanted %v", got, x)
		}
	})
}

func ExampleOr() {
	var s1 *string // nil pointer
	s2 := "hello"

	fmt.Println(pointer.Or(s1, "default"))
	fmt.Println(pointer.Or(&s2, "you won't see me"))
	// Output:
	// default
	// hello
}
