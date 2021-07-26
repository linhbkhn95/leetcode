package mergsortedarr

import (
	"reflect"
	"testing"
)

func TestMergSortedArray(t *testing.T) {
	type args struct {
		firstArr  []int
		secondArr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				firstArr:  []int{1, 2, 2, 2},
				secondArr: []int{2, 2, 2, 2, 3, 4},
			},
			want: []int{1, 2, 2, 2, 2, 2, 2, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergSortedArray(tt.args.firstArr, tt.args.secondArr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
