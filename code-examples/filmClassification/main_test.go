package main

import "testing"

func TestGetClassificationByAge(t *testing.T) {
	type args struct {
		age int
	}
	tests := []struct {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAgeClassification(tt.args.age); got != tt.want {
				t.Errorf("GetAgeClassification() = %v, want %v", got, tt.want)
			}
		})
	}
}
