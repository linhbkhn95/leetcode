package addbinary

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		// {
		// 	name: "Test 1",
		// 	args: args{
		// 		a: "11",
		// 		b: "1",
		// 	},
		// 	want: "100",
		// },
		// {
		// 	name: "Test 2",
		// 	args: args{
		// 		a: "1010",
		// 		b: "1011",
		// 	},
		// 	want: "10101",
		// },
		// {
		// 	name: "Test 3",
		// 	args: args{
		// 		a: "111",
		// 		b: "1010",
		// 	},
		// 	want: "10001",
		// },
		// {
		// 	name: "Test 4",
		// 	args: args{
		// 		a: "100",
		// 		b: "110010",
		// 	},
		// 	want: "110110",
		// },
		{
			name: "Test 5",
			args: args{
				a: "101111",
				b: "10",
			},
			want: "1010001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
