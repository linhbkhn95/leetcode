package merge_in_between_linked_lists

import (
	"reflect"
	"testing"
)

func Test_mergeInBetween(t *testing.T) {
	type args struct {
		list1 *ListNode
		a     int
		b     int
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				list1: &ListNode{
					Val: 10,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 13,
							Next: &ListNode{
								Val: 6,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 5,
									},
								},
							},
						},
					},
				},
				a: 3,
				b: 4,
				list2: &ListNode{
					Val: 1000000,
					Next: &ListNode{
						Val: 1000001,
						Next: &ListNode{
							Val:  1000002,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 13,
						Next: &ListNode{
							Val: 1000000,
							Next: &ListNode{
								Val: 1000001,
								Next: &ListNode{
									Val: 1000002,
									Next: &ListNode{
										Val: 5,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeInBetween(tt.args.list1, tt.args.a, tt.args.b, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeInBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
