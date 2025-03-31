package pointer_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/FollowTheProcess/pointer"
)

func TestNew(t *testing.T) {
	tests := []struct {
		thing any
		name  string
	}{
		{
			name:  "string",
			thing: "hello",
		},
		{
			name:  "int",
			thing: 42,
		},
		{
			name:  "bool",
			thing: true,
		},
		{
			name:  "struct",
			thing: struct{ name string }{name: "hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ptr := pointer.New(tt.thing)

			if got := reflect.TypeOf(ptr).Kind(); got != reflect.Pointer {
				t.Errorf("got %s, wanted %s", got, reflect.Pointer)
			}
		})
	}
}

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

func TestOrDefault(t *testing.T) {
	t.Run("string nil", func(t *testing.T) {
		var str *string // nil

		got := pointer.OrDefault(str)

		if got != "" {
			t.Errorf("got %q, wanted %q", got, "")
		}
	})

	t.Run("string non nil", func(t *testing.T) {
		str := "wow"

		got := pointer.OrDefault(&str)

		if got != str {
			t.Errorf("got %q, wanted %q", got, str)
		}
	})

	t.Run("int nil", func(t *testing.T) {
		var i *int // nil

		got := pointer.OrDefault(i)

		if got != 0 {
			t.Errorf("got %d, wanted %d", got, 0)
		}
	})

	t.Run("int non nil", func(t *testing.T) {
		i := 12

		got := pointer.OrDefault(&i)

		if got != i {
			t.Errorf("got %d, wanted %d", got, i)
		}
	})

	t.Run("struct nil", func(t *testing.T) {
		type thing struct{ name string } //nolint: unused // We need this
		var x *thing

		got := pointer.OrDefault(x)
		want := thing{}

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("struct non nil", func(t *testing.T) {
		type thing struct{ name string }
		x := thing{name: "x"}

		got := pointer.OrDefault(&x)

		if got != x {
			t.Errorf("got %v, wanted %v", got, x)
		}
	})
}

func ExampleNew() {
	s := pointer.New("hello")
	fmt.Printf("%T\n", s)

	// Output:
	// *string
}

func ExampleOr() {
	var s1 *string // nil pointer
	s2 := "hello"

	fmt.Printf("%q\n", pointer.Or(s1, "default"))
	fmt.Printf("%q\n", pointer.Or(&s2, "you won't see me"))

	// Output:
	// "default"
	// "hello"
}

func ExampleOrDefault() {
	var s1 *string // nil pointer
	s2 := "wow!"

	fmt.Printf("%q\n", pointer.OrDefault(s1))
	fmt.Printf("%q\n", pointer.OrDefault(&s2))

	// Output:
	// ""
	// "wow!"
}
