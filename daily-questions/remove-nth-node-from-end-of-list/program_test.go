package removenthnodefromendoflist

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
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
		// 				Val: 2,
		// 				Next: &ListNode{
		// 					Val: 3,
		// 					Next: &ListNode{
		// 						Val: 4,
		// 						Next: &ListNode{
		// 							Val: 5,
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		n: 2,
		// 	},
		// 	want: &ListNode{
		// 		Val: 1,
		// 		Next: &ListNode{
		// 			Val: 2,
		// 			Next: &ListNode{
		// 				Val: 3,
		// 				Next: &ListNode{
		// 					Val: 5,
		// 				},
		// 			},
		// 		},
		// 	},
		// },
		{
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
				n: 1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
