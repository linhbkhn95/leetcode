package stringsmatching

import (
	"reflect"
	"testing"
)

func Test_stringMatching(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// {
		// 	name: "1",
		// 	args: args{
		// 		words: []string{"mass", "as", "hero", "superhero"},
		// 	},
		// 	want: []string{"as", "hero"},
		// },
		{
			name: "2",
			args: args{
				words: []string{"blue", "green", "bu"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringMatching(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
