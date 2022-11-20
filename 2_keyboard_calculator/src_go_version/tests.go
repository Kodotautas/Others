// tests for the calculator
package main

import (
	"testing"
)

// test main function
func TestMain(_ *testing.T) {
	main()
}

// test input_variables function
func TestInput_variables(_ *testing.T) {
	input_variables()
}

// test store_keypress function
func TestStore_keypress(_ *testing.T) {
	store_keypress()
}

// test removeSpecialChars function
func TestRemoveSpecialChars(_ *testing.T) {
	removeSpecialChars(keypresses)
}

// test top_n_keypresses function
func TestTop_n_keypresses(_ *testing.T) {
	top_n_keypresses(3, keypresses, "interval")
}
