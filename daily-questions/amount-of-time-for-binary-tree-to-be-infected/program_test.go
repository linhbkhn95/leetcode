package amountoftimeforbinarytreetobeinfected

import "testing"

func Test_amountOfTime(t *testing.T) {
	type args struct {
		root  *TreeNode
		start int
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
					Val: 1,
					Left: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 9,
							},
							Right: &TreeNode{
								Val: 2,
							},
						},
					},
					Right: &TreeNode{Val: 3,
						Left:  &TreeNode{Val: 10},
						Right: &TreeNode{Val: 6}},
				},
				start: 3,
			},

			want: 4,
		},
		// {
		// 	name: "",
		// 	args: args{
		// 		root: &TreeNode{
		// 			Val: 1,
		// 			Left: &TreeNode{
		// 				Val: 2,
		// 				Left: &TreeNode{
		// 					Val: 3,
		// 					Left: &TreeNode{
		// 						Val: 4,
		// 						Left: &TreeNode{
		// 							Val: 5,
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		start: 4,
		// 	},

		// 	want: 3,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		root: &TreeNode{
		// 			Val: 1,
		// 			Right: &TreeNode{
		// 				Val: 2,
		// 				Left: &TreeNode{
		// 					Val: 3,
		// 					Right: &TreeNode{
		// 						Val: 5,
		// 					},
		// 				},
		// 				Right: &TreeNode{
		// 					Val: 4,
		// 				},
		// 			},
		// 		},
		// 		start: 4,
		// 	},

		// 	want: 3,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := amountOfTime(tt.args.root, tt.args.start); got != tt.want {
				t.Errorf("amountOfTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
