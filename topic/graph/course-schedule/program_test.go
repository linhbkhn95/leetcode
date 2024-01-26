package courseschedule

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	name: "",
		// 	args: args{
		// 		numCourses: 2,
		// 		prerequisites: [][]int{
		// 			{0, 1},
		// 		},
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		numCourses: 2,
		// 		prerequisites: [][]int{
		// 			{0, 1},
		// 			{1, 0},
		// 		},
		// 	},
		// 	want: false,
		// },
		{
			name: "",
			args: args{
				numCourses: 20,
				prerequisites: [][]int{
					{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
