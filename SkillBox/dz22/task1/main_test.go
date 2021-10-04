package main

import (
	"fmt"
	"testing"
)

func Test_findSubstr(t *testing.T) {
	type args struct {
		str    string
		substr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"T1", args{str: "строка для поиска", substr: "поиска"}, true},
		{"T2", args{str: "привет мир!", substr: "вет"}, true},
		{"T3", args{str: "Программирование - это просто", substr: "вание"}, true},
		{"T4", args{str: "Программирование - это просто", substr: "корабль"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := findSubstr(&tt.args.str, &tt.args.substr); res != tt.wantErr {
				t.Errorf("findSubstr() result = %v, need %v", res, tt.wantErr)
				fmt.Printf("%s: Fail\n", tt.name)
			} else {
				fmt.Printf("%s: Ok\n", tt.name)
			}
		})
	}
}
