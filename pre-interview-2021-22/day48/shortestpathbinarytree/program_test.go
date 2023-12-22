package program

import "testing"

func Test_getDirections(t *testing.T) {
	type args struct {
		root       *TreeNode
		startValue int
		destValue  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				startValue: 3,
				destValue:  6,
			},
			want: "UURL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDirections(tt.args.root, tt.args.startValue, tt.args.destValue); got != tt.want {
				t.Errorf("getDirections() = %v, want %v", got, tt.want)
			}
		})
	}
}
