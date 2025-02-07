package main

import "testing"

func TestFoo(t *testing.T) {
	if Foo() != 2 {
		t.Errorf("Foo() = %d; want 2", Foo())
	}
}
