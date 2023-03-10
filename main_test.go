package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := split("a/b/c", "/")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
