package main

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func Test_Is_Prime_V1(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected bool
	}{
		{
			numbers:  []int{2, 3, 5, 7, 11, 13, 17},
			expected: true,
		},
		{
			numbers:  []int{1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20},
			expected: false,
		},
		{
			numbers: []int{-1, 0, 1},
			expected: false,
		},
	}

	for _, test := range tests {
		for _, number := range test.numbers {
			result := isPrimeV1(number)
			assert.Equal(t, test.expected, result, "値が一致しません")
		}
	}
}


func Test_Is_Prime_V2(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected bool
	}{
		{
			numbers:  []int{2, 3, 5, 7, 11, 13, 17},
			expected: true,
		},
		{
			numbers:  []int{1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20},
			expected: false,
		},
		{
			numbers: []int{-1, 0, 1},
			expected: false,
		},
	}

	for _, test := range tests {
		for _, number := range test.numbers {
			result := isPrimeV2(number)

			assert.Equal(t, test.expected, result, fmt.Sprintf("値が一致しません: %d", number))
		}
	}
}

func Test_Is_Prime_V3(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected bool
	}{
		{
			numbers:  []int{2, 3, 5, 7, 11, 13, 17},
			expected: true,
		},
		{
			numbers:  []int{1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20},
			expected: false,
		},
		{
			numbers: []int{-1, 0, 1},
			expected: false,
		},
	}

	for _, test := range tests {
		for _, number := range test.numbers {
			result := isPrimeV3(number)
			assert.Equal(t, test.expected, result, "値が一致しません")
		}
	}
}

func Test_Is_Prime_V4(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected bool
	}{
		{
			numbers:  []int{2, 3, 5, 7, 11, 13, 17},
			expected: true,
		},
		{
			numbers:  []int{1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20},
			expected: false,
		},
		{
			numbers: []int{-1, 0, 1},
			expected: false,
		},
	}

	for _, test := range tests {
		for _, number := range test.numbers {
			result := isPrimeV4(number)
			assert.Equal(t, test.expected, result, fmt.Sprintf("値が一致しません: %d", number))
		}
	}
}