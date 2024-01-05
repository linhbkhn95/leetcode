package buildarrwithstackops

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	type args struct {
		target []int
		n      int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				target: []int{1, 2},
				n:      4,
			},
			want: []string{"Push", "Push", "Pop", "Push"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArray(tt.args.target, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
