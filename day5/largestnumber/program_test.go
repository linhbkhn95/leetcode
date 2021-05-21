package largestnumber

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{8308, 830},
			},
			want: "8308830",
		},
		{
			name: "test2",
			args: args{
				nums: []int{34323, 3432},
			},
			want: "343234323",
		},
		{
			name: "test3",
			args: args{
				nums: []int{0, 0},
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
