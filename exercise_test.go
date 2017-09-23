package golang

import (
	"testing"
)

func TestCaesarCipherROT4(t *testing.T) {
	original := "the quick brow fox jumps over the lazy dog"
	expected := "qeb nrfzh yolt clu grjmp lsbo qeb ixwv ald"

	s := Caesar(original, 4)

	if s != expected {
		t.Errorf("%s is expected result but got %s", expected, s)
	}
}
