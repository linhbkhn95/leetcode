package addonerowtotree

import (
	"reflect"
	"testing"
)

func Test_addOneRow(t *testing.T) {
	type args struct {
		root  *TreeNode
		val   int
		depth int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 5,
						},
					},
				},
				val:   1,
				depth: 2,
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 5,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOneRow(tt.args.root, tt.args.val, tt.args.depth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOneRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
