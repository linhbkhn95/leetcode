package imagesmoother

import (
	"reflect"
	"testing"
)

func Test_imageSmoother(t *testing.T) {
	type args struct {
		img [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				img: [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}},
			},
			want: [][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}},
		},
		{
			name: "",
			args: args{
				img: [][]int{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 10},
					{11, 12, 13, 14, 15},
					{16, 17, 18, 19, 20},
					{21, 22, 23, 24, 25}},
			},
			want: [][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imageSmoother(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageSmoother() = %v, want %v", got, tt.want)
			}
		})
	}
}
