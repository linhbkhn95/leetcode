package besttimetobuyandsellstockwithtransactionfee

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				prices: []int{1, 3, 2, 8, 4, 9},
				fee:    2,
			},
			want: 11,
		},
		// {
		// 	name: "",
		// 	args: args{
		// 		prices: []int{1, 3, 7, 5, 10, 3},
		// 		fee:    3,
		// 	},
		// 	want: 6,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
