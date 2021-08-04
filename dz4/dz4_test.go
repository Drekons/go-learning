// Параллельное исполнение
package main

import (
	"errors"
	"fmt"
	"testing"
)

func successFunc() error {
	return nil
}

func failFunc() error {
	return errors.New("testError")
}

func Test_processor(t *testing.T) {
	type args struct {
		processes    []process
		countOneTime int
		countError   int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"T1", args{processes: []process{successFunc, failFunc, failFunc}, countOneTime: 3, countError: 3}, false},
		{"T2", args{processes: []process{successFunc, failFunc, failFunc, failFunc}, countOneTime: 3, countError: 3}, true},
		{"T3", args{processes: []process{successFunc, failFunc, failFunc, successFunc, successFunc}, countOneTime: 3, countError: 3}, false},
		{"T4", args{processes: []process{successFunc, failFunc, failFunc, successFunc, successFunc}, countOneTime: 1, countError: 2}, true},
		{"T4", args{processes: []process{successFunc, failFunc, failFunc, successFunc, successFunc, successFunc, failFunc}, countOneTime: 3, countError: 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := processor(tt.args.processes, tt.args.countOneTime, tt.args.countError); (err != nil) != tt.wantErr {
				t.Errorf("processor() error = %v, wantErr %v", err, tt.wantErr)
				fmt.Printf("%s: Fail\n", tt.name)
			} else {
				fmt.Printf("%s: Ok\n", tt.name)
			}
		})
	}
}
