package internal

import (
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	cases := []struct {
		orderItems int
		packSizes  []int
		expected   map[int]int
	}{
		{1, []int{250, 500, 1000, 2000, 5000}, map[int]int{250: 1}},
		{250, []int{250, 500, 1000, 2000, 5000}, map[int]int{250: 1}},
		{251, []int{250, 500, 1000, 2000, 5000}, map[int]int{500: 1}},
		{501, []int{250, 500, 1000, 2000, 5000}, map[int]int{500: 1, 250: 1}},
		{12001, []int{250, 500, 1000, 2000, 5000}, map[int]int{5000: 2, 2000: 1, 250: 1}},

		{751, []int{250, 500, 1000, 2000, 5000}, map[int]int{1000: 1}},
		{1001, []int{250, 500, 1000, 2000, 5000}, map[int]int{1000: 1, 250: 1}},
		{1251, []int{250, 500, 1000, 2000, 5000}, map[int]int{1000: 1, 500: 1}},
		{1501, []int{250, 500, 1000, 2000, 5000}, map[int]int{1000: 1, 500: 1, 250: 1}},
		{1751, []int{250, 500, 1000, 2000, 5000}, map[int]int{2000: 1}},
		{4751, []int{250, 500, 1000, 2000, 5000}, map[int]int{5000: 1}},
		{9751, []int{250, 500, 1000, 2000, 5000}, map[int]int{5000: 2}},
		{249, []int{250, 500, 1000, 2000, 5000}, map[int]int{250: 1}},

		{3751, []int{250, 500, 1500, 2000, 4000}, map[int]int{4000: 1}},
		{14, []int{5, 12}, map[int]int{5: 3}},
		{24, []int{5, 12}, map[int]int{12: 2}},
		{500, []int{250}, map[int]int{250: 2}},
		{751, []int{5, 12}, map[int]int{5: 11, 12: 58}},
		{251, []int{5, 12, 250}, map[int]int{5: 7, 12: 18}},
		{751, []int{5, 12, 250}, map[int]int{250: 2, 12: 13, 5: 19}},
		{26, []int{5, 12}, map[int]int{12: 1, 5: 3}},
		{27, []int{5, 12}, map[int]int{12: 1, 5: 3}},
		{28, []int{5, 12}, map[int]int{12: 2, 5: 1}},
		{30, []int{5, 12}, map[int]int{5: 6}},
	}

	for _, c := range cases {
		result := CalculatePacks(c.orderItems, c.packSizes)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("calculatePacks(%d) == %v, expected %v", c.orderItems, result, c.expected)
		}
	}
}

func TestRemoveDuplicatesAndNegativeAndZeroValues(t *testing.T) {
	cases := []struct {
		packSizes  []int
		expected   []int
	}{
		{[]int{250, 500, 1000, 2000, 5000}, []int{250, 500, 1000, 2000, 5000}},
		{[]int{250, 250, 500, 500, 1000}, []int{250, 500, 1000}},
		{[]int{0, 250}, []int{250}},
		{[]int{0, 0}, make([]int, 0)},
	}

	for _, c := range cases {
		result := removeDuplicatesAndNegativeAndZeroValues(c.packSizes)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("removeDuplicatesAndNegativeAndZeroValues(%d) == %v, expected %v", c.packSizes, result, c.expected)
		}
	}
}
