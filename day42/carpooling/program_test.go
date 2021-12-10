package carpooling

import "testing"

func Test_carPooling(t *testing.T) {
	type args struct {
		trips    [][]int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		// {
		// 	name: "Test 1",
		// 	args: args{
		// 		trips:    [][]int{{2, 1, 5}, {3, 5, 7}},
		// 		capacity: 3,
		// 	},
		// 	want: true,
		// },
		{
			name: "Test 1",
			args: args{
				trips:    [][]int{{8, 2, 3}, {4, 1, 3}, {1, 3, 6}, {8, 4, 6}, {4, 4, 8}},
				capacity: 12,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carPooling(tt.args.trips, tt.args.capacity); got != tt.want {
				t.Errorf("carPooling() = %v, want %v", got, tt.want)
			}
		})
	}
}
