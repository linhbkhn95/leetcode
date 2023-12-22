package removeduplicate

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// {
		// 	name: "test 1",
		// 	args: args{
		// 		head: &ListNode{
		// 			Val: 1,
		// 			Next: &ListNode{
		// 				Val: 2,
		// 				Next: &ListNode{
		// 					Val: 2,
		// 					Next: &ListNode{
		// 						Val: 3,
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// 	want: &ListNode{
		// 		Val: 1,
		// 		Next: &ListNode{
		// 			Val: 3,
		// 		},
		// 	},
		// },

		{
			name: "test 2",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 4,
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
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreeSum(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test 1",
			args: args{
				arr: []int{-1, 0, 1, 2, -1, -4},
				k:   0,
			},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSum(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
