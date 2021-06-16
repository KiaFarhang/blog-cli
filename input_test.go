package main

import "testing"

func TestIsBlankString(t *testing.T) {
	assert := func(t testing.TB, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	}

	t.Run("empty string", func(t *testing.T) {
		assert(t, isBlankString(""), true)
	})

	t.Run("whitespace string", func(t *testing.T) {
		assert(t, isBlankString("    "), true)
	})

	t.Run("non-empty string", func(t *testing.T) {
		assert(t, isBlankString("k"), false)
	})
}
