package main

import (
	"testing"
)

func Test_unpackStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"a4bc2d5e", args{str: "a4bc2d5e"}, "aaaabccddddde"},
		{"abcd", args{str: "abcd"}, "abcd"},
		{"45", args{str: "45"}, ""},
		{`qwe\4\5`, args{str: `qwe\4\5`}, "qwe45"},
		{`qwe\45`, args{str: `qwe\45`}, "qwe44444"},
		{`qwe\\5`, args{str: `qwe\\5`}, `qwe\\\\\`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unpackStr(tt.args.str); got != tt.want {
				t.Errorf("unpackStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reapetedString(t *testing.T) {
	type args struct {
		str   string
		count int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"a3", args{str: "a", count: 3}, "aaa"},
		{"ab2", args{str: "ab", count: 2}, "abab"},
		{"abc0", args{str: "abc", count: 0}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reapetedString(tt.args.str, tt.args.count); got != tt.want {
				t.Errorf("reapetedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
