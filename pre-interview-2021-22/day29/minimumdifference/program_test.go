package program

import (
	"reflect"
	"testing"
)

func Test_minimumAbsDifference(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test 1",
			args: args{
				arr: []int{3, 8, -10, 23, 19, -4, -14, 27},
			},
			want: [][]int{{-14, -10}, {19, 23}, {23, 27}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAbsDifference(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumAbsDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
