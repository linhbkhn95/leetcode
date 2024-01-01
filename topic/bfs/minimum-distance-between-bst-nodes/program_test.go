package minimumdistancebetweenbstnodes

import "testing"

func Test_minDiffInBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 90,
					Left: &TreeNode{
						Val: 69,
						Left: &TreeNode{
							Val: 49,
							Right: &TreeNode{
								Val: 52,
							},
						},
						Right: &TreeNode{
							Val: 89,
						},
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDiffInBST(tt.args.root); got != tt.want {
				t.Errorf("minDiffInBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
