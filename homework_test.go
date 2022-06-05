package reverse_string

import (
	"testing"
)

const MESSAGE = "%+v => %+v"
const MESSAGE_LOG = "\n%+35v => %+35v\n%+74v"

func TestReverseString(t *testing.T) {
	in_string := "Hello, Go language /ˈlæŋ.ɡwɪdʒ/"
	got := ReverseString(in_string)
	want := "/ʒdɪwɡ.ŋælˈ/ egaugnal oG ,olleH"
	t.Logf(MESSAGE_LOG, in_string, got, want)

	if got != want {
		t.Errorf(MESSAGE, got, want)
	}
}
