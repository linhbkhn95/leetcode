package partitionlist

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// {
		// 	name: "",
		// 	args: args{
		// 		head: &ListNode{
		// 			Val: 1,
		// 			Next: &ListNode{
		// 				Val: 4,
		// 				Next: &ListNode{
		// 					Val: 3,
		// 					Next: &ListNode{
		// 						Val: 2,
		// 						Next: &ListNode{
		// 							Val: 5,
		// 							Next: &ListNode{
		// 								Val: 2,
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		x: 3,
		// 	},
		// 	want: &ListNode{
		// 		Val: 1,
		// 		Next: &ListNode{
		// 			Val: 2,
		// 			Next: &ListNode{
		// 				Val: 2,
		// 				Next: &ListNode{
		// 					Val: 4,
		// 					Next: &ListNode{
		// 						Val: 3,
		// 						Next: &ListNode{
		// 							Val: 5,
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// },
		{
			name: "",
			args: args{
				head: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
				x: 2,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
