package day12

import (
	"reflect"
	"testing"
)

func TestHashify(t *testing.T) {
	type args struct {
		strLen int
	}
	tests := []struct {
		name string
		args args
		want [][]rune
	}{
		{
			name: "3 long",
			args: args{strLen: 3},
			want: [][]rune{
				[]rune("..."),
				[]rune("#.."),
				[]rune(".#."),
				[]rune("..#"),
				[]rune("##."),
				[]rune(".##"),
				[]rune("#.#"),
				[]rune("###"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hashify(tt.args.strLen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hashify() = %v, want %v", got, tt.want)
			}
		})
	}
}
