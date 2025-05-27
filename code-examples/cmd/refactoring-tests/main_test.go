package main

import "testing"

//exportable using capitals - uppercase means it can read, lowercase it cannot
//go test main_test.go
func TestItworks(t *testing.T) {
	passed := true
	if !passed {
		t.Error("test failed but testing works")
	}
}

// //f12 to test how functions are working
// func TestGreeting_Evening21(t *testing.T) {
// 	time := 21 //within function you dont have to statically type, allowing the compile to assume the type (like below)
// 	expected := "Good evening"
// 	actual := greeting(time)
// 	if actual != expected {
// 		t.Errorf(`expected "%v" expected, but got "%v"`, expected, actual) //Errorf for formatting
// 	}
// }

// func TestGreeting_Evening4(t *testing.T) {
// 	time := 4
// 	expected := "Good evening"
// 	actual := greeting(time)
// 	if actual != expected {
// 		t.Errorf(`expected "%v" expected, but got "%v"`, expected, actual)
// 	}
// }

// func TestGreeting_Morning8(t *testing.T) {
// 	time := 8
// 	expected := "Good morning"
// 	actual := greeting(time)
// 	if actual != expected {
// 		t.Errorf(`expected "%v" expected, but got "%v"`, expected, actual)
// 	}
// }

// func TestGreeting_Afternoon16(t *testing.T) {
// 	time := 16
// 	expected := "Good afternoon"
// 	actual := greeting(time)
// 	if actual != expected {
// 		t.Errorf(`expected "%v" expected, but got "%v"`, expected, actual)
// 	}
// }

// func TestGreeting(t *testing.T) {
// 	tests := []struct {
// 		time     int
// 		expected string
// 	}{
// 		{time: 4, expected: "Good evening"},
// 		{time: 8, expected: "Good morning"},
// 		{time: 16, expected: "Good afternoon"},
// 		{time: 21, expected: "Good evening"},
// 	}

// 	for _, test := range tests {	//_ means variable is there but we are not using it - range controls index and element
// 		actual := greeting(test.time)
// 		if actual != test.expected {
// 			t.Errorf("For time %d, expected %q but got %q", test.time, test.expected, actual)
// 		}
// 	}
// }

func TestGreeting_Morning(t *testing.T) {
	tests := []int{5, 11} //this is a slice (dynamic array/list)
	expected := "Good morning"

	for _, tc := range tests {
		t.Run("Test Morning", func(t *testing.T) {
			actual := greeting(tc)
			if actual != expected {
				t.Errorf(`expected "%v" but got "%v"`, expected, actual)
			}
		})
	}
}

func TestGreeting_BoundaryCases(t *testing.T) {
	tests := []struct {
		name     string
		time     int
		greeting string
	}{
		{"Test Evening", 21, "Good evening"},
		{"Test Morning", 5, "Good morning"},
		{"Test Afternoon", 15, "Good afternoon"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := greeting(tc.time)
			if actual != tc.greeting {
				t.Errorf(`expected "%v" but got "%v"`, tc.greeting, actual)
			}
		})
	}
}

// equivalence partitioning is a tactic used to cut down the amount of test cases you need when unit testing
// divides the possible inputs into categories
