package main

import "testing"

func Test_pseudoPalindromicPaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 1,
							},
						},
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pseudoPalindromicPaths(tt.args.root); got != tt.want {
				t.Errorf("pseudoPalindromicPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
