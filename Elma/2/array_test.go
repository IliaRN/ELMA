package main

import "testing"

func TestSolution(t *testing.T) {
	success1 := false
	case1 := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	result := Solution(case1, k)
	expectedFor1 := []int{5, 6, 7, 1, 2, 3, 4}

	for i := range result {
		if result[i] != expectedFor1[i] {
			success1 = false
			break
		}
		success1 = true
	}
	if !success1 {
		t.Errorf("Wrong result. Expected: %v, Got: %v ", expectedFor1, result)
	}

	success2 := false
	case2 := []int{-1, 0, 0, 0, 3, 1, 99, -99, 32, 1}
	k = 10

	result = Solution(case2, k)
	expectedFor2 := []int{-1, 0, 0, 0, 3, 1, 99, -99, 32, 1}

	for i := range result {
		if result[i] != expectedFor2[i] {
			success2 = false
			break
		}
		success2 = true
	}
	if !success2 {
		t.Errorf("Wrong result. Expected: %v, Got: %v ", expectedFor2, result)
	}

	// Negative case
	success3 := false
	case3 := []int{-1, 0, 0, 0, 3, 1, 99, -99, 32, 1}
	k = 10

	result = Solution(case3, k)
	expectedFor3 := []int{-1, 0, 0, 0, 3, 1, 99, -99, 32, 2}

	for i := range result {
		if result[i] != expectedFor3[i] {
			success2 = false
			break
		}
		success3 = true
	}
	if !success3 {
		t.Errorf("Wrong result. Expected: %v, Got: %v ", expectedFor2, result)
	}
}

func TestSecondSolution(t *testing.T) {
	success1 := false
	case1 := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	result := SecondSolution(case1, k)
	expectedFor1 := []int{5, 6, 7, 1, 2, 3, 4}

	for i := range result {
		if result[i] != expectedFor1[i] {
			success1 = false
			break
		}
		success1 = true
	}
	if !success1 {
		t.Errorf("Wrong result. Expected: %v, Got: %v ", expectedFor1, result)
	}

	success2 := false
	case2 := []int{-1, 0, 0, 0, 3, 1, 99, -99, 32, 1}
	k = 10

	result = SecondSolution(case2, k)
	expectedFor2 := []int{-1, 0, 0, 0, 3, 1, 99, -99, 32, 1}

	for i := range result {
		if result[i] != expectedFor2[i] {
			success2 = false
			break
		}
		success2 = true
	}
	if !success2 {
		t.Errorf("Wrong result. Expected: %v, Got: %v ", expectedFor2, result)
	}

	// Negative case
	success3 := false
	case3 := []int{-1, 0, 0, 0, 3, 1, 99, -99, 32, 1}
	k = 10

	result = SecondSolution(case3, k)
	expectedFor3 := []int{-1, 0, 0, 0, 3, 1, 99, -99, 32, 2}

	for i := range result {
		if result[i] != expectedFor3[i] {
			success2 = false
			break
		}
		success3 = true
	}
	if !success3 {
		t.Errorf("Wrong result. Expected: %v, Got: %v ", expectedFor2, result)
	}
}

