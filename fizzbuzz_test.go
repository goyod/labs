package main_test

import(
	"testing"
	
	"github.com/goyod/labs"
)

func TestFizzBuzzGivenOneReturnsOne(t *testing.T) {
	given := 1

	want := "1"

	get := main.FizzBuzz(given)

	if want != get {
		t.Errorf("given %d wants %q but get %q\n",given,want,get)
	}
}