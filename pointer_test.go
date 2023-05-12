package pointer_test

import (
	"testing"

	"github.com/FollowTheProcess/pointer"
)

func TestHello(t *testing.T) {
	got := pointer.Hello()
	want := "Hello pointer"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
