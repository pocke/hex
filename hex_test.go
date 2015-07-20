package main

import "testing"

func TestHex(t *testing.T) {
	assert := func(dec, hex string) {
		if v, err := Hex(dec); err != nil {
			t.Error(err)
		} else if v != hex {
			t.Errorf("Expected: %s, but got %s", hex, v)
		}
	}

	assert("10", "A")
	assert("255", "FF")
	assert("256", "100")
	assert("3", "3")
}

func TestDec(t *testing.T) {
	assert := func(hex, dec string) {
		if v, err := Dec(hex); err != nil {
			t.Error(err)
		} else if v != dec {
			t.Errorf("Expected: %s, but got %s", dec, v)
		}
	}

	assert("A", "10")
	assert("FF", "255")
	assert("100", "256")
	assert("3", "3")
}
