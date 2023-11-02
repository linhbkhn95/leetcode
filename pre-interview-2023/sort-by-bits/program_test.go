package sortbybits

import (
	"reflect"
	"testing"
)

func Test_sortByBits(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				arr: []int{1111, 7644, 1107, 6978, 8742, 1, 7403, 7694, 9193, 4401, 377, 8641, 5311, 624, 3554, 6631},
			},
			want: []int{1, 624, 1107, 4401, 8641, 8742, 377, 1111, 6978, 3554, 7694, 9193, 5311, 6631, 7403, 7644},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByBits(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
