package main

import (
	"testing"
)

func TestMedian(t *testing.T) {
	tests := []struct {
		numders  []int
		expected float64
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0},
		{[]int{5, 3, 1, 4, 2}, 3.0}, // Неотсортированный вход
		{[]int{-1, -2, -3}, -2.0},
		{[]int{1, 2}, 1.5},
		{[]int{100, 200, 300}, 200.0},
		{[]int{1}, 1.0},
		{[]int{}, 0.0}, // Пустой срез, нужно обработать отдельно
		{[]int{10, 10, 10}, 10.0},
		{[]int{1, 2, 3, 4, 5, 6}, 3.5},
		{[]int{-1, 0, 1}, 0.0},
	}

	for _, test := range tests {
		result := Median(test.numders)
		if result != test.expected {
			t.Errorf("Median(%v) result - %.2f expected - %.2f ", test.numders, result, test.expected)
		}
	}
}

func TestMean(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected float64
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0},
		{[]int{-1, 0, 1}, 0.0},
		{[]int{-1, -2, -3}, -2.0},
		{[]int{100, 200, 300}, 200.0},
		{[]int{}, 0.0},
		{[]int{10, 10, 10, 10}, 10.0},
		{[]int{1}, 1.0},
		{[]int{-100000, 100000}, 0.0},
		{[]int{1000, 2000, 3000, 4000}, 2500.0},
		{[]int{1, 2, 3, 4, 5, 6}, 3.5},
	}
	for _, test := range tests {
		result := Mean(test.numbers)
		if result != test.expected {
			t.Errorf("Mean(%v) result - %.2f expected - %.2f", test.numbers, result, test.expected)
		}
	}
}

func TestMode(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected int
	}{
		{[]int{1, 2, 2, 3}, 2},
		{[]int{1, 1, 2, 2}, 1}, // Наименьшая мода
		{[]int{-1, -1, -2, -3}, -1},
		{[]int{5, 5, 5, 2, 3, 3}, 5},
		{[]int{10, 20, 20, 30, 30}, 20}, // Наименьшая мода
		{[]int{}, 0},                    // Пустой срез, нужно обработать отдельно
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 1},          // Все уникальные, возвращаем первое
		{[]int{1, 1, 2, 2, 3}, 1},    // Наименьшая мода среди нескольких
		{[]int{2, 3, 4, 2, 3, 3}, 3}, // Проверка на частоту
	}

	for _, test := range tests {
		result := Mode(test.numbers)
		if result != test.expected {
			t.Errorf("Mode(%v) = %d; expected %d", test.numbers, result, test.expected)
		}
	}
}

func TestSd(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected float64
	}{
		{[]int{-1, 0, 1}, 1.00},
		{[]int{100, 200, 300}, 100.00},
		{[]int{-100000, 100000}, 141421.36},
		{[]int{1, 2, 3, 4, 5, 6}, 1.87},
		{[]int{-10, 0, 10}, 10.00},
		{[]int{2, 4, 4, 4, 5, 5, 7, 9}, 2.14},
		{[]int{1, 1, 1, 1, 2}, 0.45},
		{[]int{}, 0.0},
		{[]int{5}, 0.0},
		{[]int{5, 5, 5, 5}, 0.0},
	}

	for _, test := range tests {
		mean := Mean(test.numbers)
		result := Sd(test.numbers, mean)
		if result < test.expected-0.01 || result > test.expected+0.01 {
			t.Errorf("Sd(%v) result - %.2f expected - %.2f", test.numbers, result, test.expected)
		}
	}
}
