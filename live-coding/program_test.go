package livecoding

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		emps []*Employee
		id   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				emps: []*Employee{
					{
						ID:         1,
						Importance: 5,
						Subs:       []int{2, 3},
					},
					{
						ID:         2,
						Importance: 3,
						Subs:       []int{},
					},
					{
						ID:         3,
						Importance: 3,
						Subs:       []int{},
					},
				},
				id: 1,
			},
			want: 11,
		},
		{
			name: "",
			args: args{
				emps: []*Employee{
					{
						ID:         1,
						Importance: 7,
						Subs:       []int{2, 3},
					},
					{
						ID:         2,
						Importance: 1,
						Subs:       []int{},
					},
					{
						ID:         3,
						Importance: 2,
						Subs:       []int{4},
					},
					{
						ID:         4,
						Importance: 3,
						Subs:       []int{},
					},
				},
				id: 1,
			},
			want: 13,
		},
		{
			name: "",
			args: args{
				emps: []*Employee{
					{
						ID:         1,
						Importance: 5,
						Subs:       []int{2, 3},
					},
					{
						ID:         2,
						Importance: 3,
						Subs:       []int{4},
					},
					{
						ID:         3,
						Importance: 3,
						Subs:       []int{4},
					},
					{
						ID:         4,
						Importance: 2,
						Subs:       []int{},
					},
				},
				id: 1,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.emps, tt.args.id); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
