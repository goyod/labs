package main_test

import (
	"testing"

	"github.com/goyod/labs"
)

func TestFizzBuzzReturnsOriginNumber(t *testing.T) {
	given := []int{1, 2, 4, 7, 8, 11, 13, 14, 16, 17, 19}
	wants := []string{"1", "2", "4", "7", "8", "11", "13", "14", "16", "17", "19"}

	for i, want := range wants {
		get := main.FizzBuzz(given[i])

		if want != get {
			t.Errorf("given %d wants %q but get %q\n", given, want, get)
		}
	}
}

func TestFizzBuzzReturnsFizzForMultiplesOfThree(t *testing.T) {
	given := []int{3, 6, 9, 12, 18}
	wants := "Fizz"

	for _, give := range given {
		get := main.FizzBuzz(give)

		if wants != get {
			t.Errorf("given %d wants %q but get %q\n", give, wants, get)
		}
	}
}

func TestFizzBuzzReturnsBuzzForMultiplesOfFive(t *testing.T) {
	given := []int{5, 10, 20, 25}
	wants := "Buzz"

	for _, give := range given {
		get := main.FizzBuzz(give)

		if wants != get {
			t.Errorf("given %d wants %q but get %q\n", give, wants, get)
		}
	}
}

func TestFizzBuzzReturnsFizzBuzzForMultiplesOfFifteen(t *testing.T) {
	given := []int{15, 30, 45, 60, 75}
	wants := "FizzBuzz"

	for _, give := range given {
		get := main.FizzBuzz(give)

		if wants != get {
			t.Errorf("given %d wants %q but get %q\n", give, wants, get)
		}
	}
}
