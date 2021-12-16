package lightbulbcircuit

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		states []int32
	}
	tests := []struct {
		name       string
		args       args
		wantResult int32
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				states: []int32{2, 1, 3, 5, 4},
			},
			wantResult: 3,
		},
		{
			name: "test 2",
			args: args{
				states: []int32{5, 4, 3, 2, 1},
			},
			wantResult: 1,
		},
		{
			name: "test 3",
			args: args{
				states: []int32{11, 1, 2, 3, 4, 7, 6, 5, 8, 9, 10},
			},
			wantResult: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Solution(tt.args.states); gotResult != tt.wantResult {
				t.Errorf("Solution() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
