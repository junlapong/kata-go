package main

import "testing"

func TestHello(t *testing.T) {
	assert := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		name := "Jun"
		got := Hello(name)
		want := "Hello " + name

		assert(t, got, want)
	})
}
