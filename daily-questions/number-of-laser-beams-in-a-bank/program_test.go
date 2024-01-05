package numberoflaserbeamsinabank

import "testing"

func Test_numberOfBeams(t *testing.T) {
	type args struct {
		bank []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				bank: []string{"011001", "000000", "010100", "001000"},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfBeams(tt.args.bank); got != tt.want {
				t.Errorf("numberOfBeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
