package findallgroupsoffarmland

import (
	"reflect"
	"testing"
)

func Test_findFarmland(t *testing.T) {
	type args struct {
		land [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				land: [][]int{
					{1, 0, 0}, {0, 1, 1}, {0, 1, 1},
				},
			},
			want: [][]int{
				{0, 0, 0, 0}, {1, 1, 2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFarmland(tt.args.land); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFarmland() = %v, want %v", got, tt.want)
			}
		})
	}
}
