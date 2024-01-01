package recoverbinarysearchtree

import "testing"

func Test_recoverTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		// {
		// 	name: "",
		// 	args: args{
		// 		root: &TreeNode{
		// 			Val: 1,
		// 			Left: &TreeNode{
		// 				Val:   3,
		// 				Right: &TreeNode{Val: 2},
		// 			},
		// 		},
		// 	},
		// },
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.args.root)
		})
	}
}
