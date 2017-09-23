package golang

import (
	"testing"

	"github.com/pallat/golang/public"
)

func TestMultipleFunc1(t *testing.T) {
	result := public.Multiple(2, 3)

	if result != 6 {
		t.Error("Please make the public multiple function")
	}
}

func TestMultipleFunc2(t *testing.T) {
	result := public.Multiple(3, 4)

	if result != 12 {
		t.Error("Please make the public multiple function")
	}
}
