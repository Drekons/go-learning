package main

import (
	"testing"
)

func Test_in_array(t *testing.T) {
	type args struct {
		element string
		array   []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test", args{element: "test", array: []string{"Hello", "world", "test"}}, true},
		{"world", args{element: "world", array: []string{"Hello", "world", "test"}}, true},
		{"fail", args{element: "worasdld", array: []string{"Hello", "world", "test"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := in_array(tt.args.element, tt.args.array); got != tt.want {
				t.Errorf("in_array() = %v, want %v", got, tt.want)
			}
		})
	}
}
