 package main

// import (
// 	"reflect"
// 	"testing"
// )

// func TestCountBits(t *testing.T) {
// 	type args struct {
// 		num int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []int
// 	}{
// 		{
// 			name: "test 1",
// 			args: args{
// 				num: 2,
// 			},
// 			want: []int{0, 1, 1},
// 		},
// 		{
// 			name: "test 1",
// 			args: args{
// 				num: 5,
// 			},
// 			want: []int{0, 1, 1, 2, 1, 2},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := CountBits(tt.args.num); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("CountBits() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_countPrimes(t *testing.T) {
// 	type args struct {
// 		n int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// {
// 		// 	name: "test 1",
// 		// 	args: args{
// 		// 		n: 10,
// 		// 	},
// 		// 	want: 4,
// 		// },
// 		// {
// 		// 	name: "test 2",
// 		// 	args: args{
// 		// 		n: 2,
// 		// 	},
// 		// 	want: 1,
// 		// },
// 		// {
// 		// 	name: "test 1",
// 		// 	args: args{
// 		// 		n: 10,
// 		// 	},
// 		// 	want: 4,
// 		// },
// 		{
// 			name: "test 2",
// 			args: args{
// 				n: 5 * 1000000,
// 			},
// 			want: 348513,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := countPrimes(tt.args.n); got != tt.want {
// 				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_gcdOfStrings(t *testing.T) {
// 	type args struct {
// 		str1 string
// 		str2 string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "test1",
// 			args: args{
// 				str1: "ABCABC",
// 				str2: "ABC",
// 			},
// 			want: "ABCc",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := gcdOfStrings(tt.args.str1, tt.args.str2); got != tt.want {
// 				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestSmallestPrimeFactor(t *testing.T) {
// 	type args struct {
// 		n int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "test1",
// 			args: args{
// 				n: 1,
// 			},
// 			want: 1,
// 		},
// 		{
// 			name: "test2",
// 			args: args{
// 				n: 2,
// 			},
// 			want: 2,
// 		},
// 		{
// 			name: "test3",
// 			args: args{
// 				n: 3,
// 			},
// 			want: 3,
// 		},
// 		{
// 			name: "test4",
// 			args: args{
// 				n: 4,
// 			},
// 			want: 2,
// 		},
// 		{
// 			name: "test5",
// 			args: args{
// 				n: 6,
// 			},
// 			want: 2,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := SmallestPrimeFactor(tt.args.n); got != tt.want {
// 				t.Errorf("SmallestPrimeFactor() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
