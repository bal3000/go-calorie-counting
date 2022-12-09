package main

import (
	"testing"
)

func TestLoadReindeerCalories(t *testing.T) {
	results, err := loadReindeerCalories("./mock.txt")
	if err != nil {
		t.Error(err)
	}

	expected := [][]int{
		{7844, 1892, 10928, 4800, 9730, 3485, 7293},
		{11956, 2767, 12893, 2757, 3026, 9979},
		{4507, 4199, 2855, 1792, 2041, 4162, 3855, 2650, 2165, 5464, 2125, 4979, 4878, 6301},
	}

	for i, r := range results {
		for j, v := range r {
			if v != expected[i][j] {
				t.Errorf("expected %d but got %d\n", expected[i][j], v)
			}
		}
	}
}

func TestCalculateTotals(t *testing.T) {
	input := [][]int{
		{7844, 1892, 10928, 4800, 9730, 3485, 7293},
		{11956, 2767, 12893, 2757, 3026, 9979},
		{4507, 4199, 2855, 1792, 2041, 4162, 3855, 2650, 2165, 5464, 2125, 4979, 4878, 6301},
	}

	expected := []int{45972, 43378, 51973}

	result := calculateTotals(input)

	for i, r := range result {
		e := expected[i]
		if r != e {
			t.Errorf("expected %d but got %d\n", e, r)
		}
	}
}

func TestSortHighestCalories(t *testing.T) {
	input := []int{45972, 43378, 51973}
	expected := []int{51973, 45972, 43378}

	result := sortHighestCalories(input)
	for i, r := range result {
		e := expected[i]
		if r != e {
			t.Errorf("expected %d but got %d\n", e, r)
		}
	}
}

func TestSumOfCalories(t *testing.T) {
	input := []int{51973, 45972, 43378}
	expected := 141323

	result := sumOfCalories(input...)
	if expected != result {
		t.Errorf("expected %d but got %d\n", expected, result)
	}
}
