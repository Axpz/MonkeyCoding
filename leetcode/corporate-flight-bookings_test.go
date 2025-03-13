package leetcode

import (
	"reflect"
	"testing"
)

func TestCorpFlightBookings(t *testing.T) {
	bookings := [][]int{
		{1, 2, 10},
		{2, 3, 20},
		{2, 5, 25},
	}
	n := 5
	expected := []int{10, 55, 45, 25, 25}

	result := corpFlightBookings(bookings, n)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
