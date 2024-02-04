package main

import (
	"reflect"
	"testing"
)

func Test_sequentialDigits(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				low:  58,
				high: 155,
			},
			want: []int{67, 78, 89, 123},
		},
		{
			name: "",
			args: args{
				low:  100,
				high: 300,
			},
			want: []int{123, 234},
		},
		{
			name: "",
			args: args{
				low:  1000,
				high: 13000,
			},
			want: []int{1234, 2345, 3456, 4567, 5678, 6789, 12345},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sequentialDigits(tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sequentialDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
