package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test_1",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			}},
			want: []int{1, 3, 5, 2, 4},
		},

		{
			name: "test_2",
			args: args{head: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val:  7,
										Next: nil,
									},
								},
							},
						},
					},
				},
			}},
			want: []int{2, 3, 6, 7, 1, 5, 4},
		},
		{
			name: "test_3",
			args: args{head: nil},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := oddEvenList(tt.args.head)
			arr := make([]int, 0)
			for got != nil {
				arr = append(arr, got.Val)
				got = got.Next
			}

			require.Equal(t, tt.want, arr)
		})
	}
}
