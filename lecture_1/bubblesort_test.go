package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {

	testCases := []struct {
		Name         string
		In           []int
		Out          []int
		Error        bool
		ErrorMessage string
	}{
		{
			Name:         "Exactly 5 items",
			In:           []int{1, 3, -1, 2, 9},
			Out:          []int{-1, 1, 2, 3, 9},
			Error:        false,
			ErrorMessage: "",
		}, {
			Name:         "Less than 5 items",
			In:           []int{1, 2, 3, 4},
			Out:          nil,
			Error:        true,
			ErrorMessage: "can sort exactly 5 numbers only",
		}, {
			Name:         "More than 5 items",
			In:           []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			Out:          nil,
			Error:        true,
			ErrorMessage: "can sort exactly 5 numbers only",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			res, err := BubbleSort(tc.In...)

			if err != nil && !tc.Error {
				t.Errorf("Error not expected! Trace: %s", err)
			}

			assert.Equal(t, res, tc.Out)
		})
	}

}
