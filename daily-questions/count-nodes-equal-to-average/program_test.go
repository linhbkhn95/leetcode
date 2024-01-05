package countnodesequaltoaverage

import "testing"

func Test_averageOfSubtree(t *testing.T) {
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
					Val: 4,
					Left: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfSubtree(tt.args.root); got != tt.want {
				t.Errorf("averageOfSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
