package checkifarraypairsaredivisiblebyk

import "testing"

func Test_canArrange(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				arr: []int{
					1, 2, 3, 4, 5, 10, 6, 7, 8, 9,
				},
				k: 5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canArrange(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("canArrange() = %v, want %v", got, tt.want)
			}
		})
	}
}
