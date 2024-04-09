package timeneededtobuytickets

import "testing"

func Test_timeRequiredToBuy(t *testing.T) {
	type args struct {
		tickets []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				tickets: []int{
					2, 3, 2,
				},
				k: 2,
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				tickets: []int{
					5, 1, 1, 1,
				},
				k: 0,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeRequiredToBuy(tt.args.tickets, tt.args.k); got != tt.want {
				t.Errorf("timeRequiredToBuy() = %v, want %v", got, tt.want)
			}
		})
	}
}
