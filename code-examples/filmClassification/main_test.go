package main

import "testing"

func TestGetClassificationByAge(t *testing.T) {
	//args is scalable - helps if function has multiple input parameters - easier to organise
	type args struct { //makes test cases
		age int
	}
	tests := []struct { //slice of anonymous structs, where each represents a test case
		name string
		args args
		want string
	}{
		{"age3-U+PG", args{3}, "U & PG films are available"},
		{"age12-U+PG+12", args{12}, "U, PG & 12 films are available"},
		{"age14-U+PG+12", args{14}, "U, PG & 12 films are available"},
		{"age15-U+PG+12+15", args{15}, "U, PG, 12 & 15 films are available"},
		{"age17-U+PG+12+15", args{17}, "U, PG, 12 & 15 films are available"},
		{"age18-All", args{18}, "All films are available"},
	}
	//tt is current test case, _ index so is unused
	for _, tt := range tests { //runs test cases - runs each element in tests
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAgeClassification(tt.args.age); got != tt.want {
				t.Errorf("GetAgeClassification() = %v, want %v", got, tt.want)
			}
		})
	}
}
