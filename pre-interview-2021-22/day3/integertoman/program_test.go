
package integertoman

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{ 
				num: 1994,
			},
			want: "MCMXCIV",
		},
		{
			name: "test2",
			args: args{ 
				num: 3,
			},
			want: "III",
		},
		{
			name: "test3",
			args: args{ 
				num: 14,
			},
			want: "XIV",
		},
		{
			name: "test3",
			args: args{ 
				num: 1000,
			},
			want: "M",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
