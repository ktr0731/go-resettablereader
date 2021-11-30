package resettablereader

import (
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	t.Parallel()

	r := New(strings.NewReader("abcdefg"))

	b := make([]byte, 2)
	n, err := r.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 2, n; want != got {
		t.Errorf("want %d, but got %d", 2, n)
	}
	if want, got := "ab", string(b); want != got {
		t.Errorf("want '%s', but got '%s'", want, got)
	}

	b = make([]byte, 2)
	n, err = r.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 2, n; want != got {
		t.Errorf("want %d, but got %d", 2, n)
	}
	if want, got := "cd", string(b); want != got {
		t.Errorf("want '%s', but got '%s'", want, got)
	}

	r.Reset()

	b = make([]byte, 2)
	n, err = r.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 2, n; want != got {
		t.Errorf("want %d, but got %d", 2, n)
	}
	if want, got := "ab", string(b); want != got {
		t.Errorf("want '%s', but got '%s'", want, got)
	}
}
