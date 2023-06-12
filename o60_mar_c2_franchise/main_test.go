package main

import (
	"reflect"
	"testing"
)

func TestSomeFn(t *testing.T) {
	type testcase struct {
		casename                       string
		roomLength, roomScope, roomMax int
		roomProfit                     []int
		expectProfit                   int
		expectArr                      []int
	}

	testcases := []testcase{
		{
			casename:     "case 1",
			roomLength:   6,
			roomScope:    4,
			roomMax:      2,
			roomProfit:   []int{1, 2, 3, 4, 5, 6},
			expectProfit: 14,
			expectArr:    []int{0, 1, 4, 5},
		},
		{
			casename:     "case 2",
			roomLength:   6,
			roomScope:    4,
			roomMax:      2,
			roomProfit:   []int{1, 4, 3, 5, 2, 6},
			expectProfit: 15,
			expectArr:    []int{1, 3, 5},
		},
		{
			casename:     "case 3",
			roomLength:   8,
			roomScope:    4,
			roomMax:      3,
			roomProfit:   []int{1, 1, 1, 2, 3, 2, 1, 1},
			expectProfit: 10,
			expectArr:    []int{2, 6},
		},
	}

	for _, tc := range testcases {
		actualProfit, actualArr := someFn(tc.roomLength, tc.roomScope, tc.roomMax, tc.roomProfit)

		if !reflect.DeepEqual(actualProfit, tc.expectProfit) {
			t.Errorf("Expected %v, but got %v", tc.expectProfit, actualProfit)
		}

		if !reflect.DeepEqual(actualArr, tc.expectArr) {
			t.Errorf("Expected %v, but got %v", tc.expectArr, actualArr)
		}
	}
}

func TestSliceEveryPossible(t *testing.T) {
	type testcase struct {
		casename         string
		roomScope        int
		roomProfit       []int
		expectSlicedRoom [][][]int
	}

	testcases := []testcase{
		{
			casename:   "case 1",
			roomScope:  4,
			roomProfit: []int{1, 2, 3, 4, 5, 6},
			expectSlicedRoom: [][][]int{
				{
					{1, 2, 3, 4},
					{5, 6},
				},
				{
					{1},
					{2, 3, 4, 5},
					{6},
				},
				{
					{1, 2},
					{3, 4, 5, 6},
				},
			},
		},
	}

	for _, tc := range testcases {
		actualslicedRoom := sliceEveryPossible(tc.roomScope, tc.roomProfit)

		if !reflect.DeepEqual(actualslicedRoom, tc.expectSlicedRoom) {
			t.Errorf("Expected %v, but got %v", tc.expectSlicedRoom, actualslicedRoom)
		}
	}
}

// func TestGetMaxByLimit

// func TestSumProfit
