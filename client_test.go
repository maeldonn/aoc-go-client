package aocgoclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzleExists(t *testing.T) {
	type args struct {
		year, day int
	}

	tests := map[string]struct {
		args args
		want bool
	}{
		"invalid year and day": {
			args: args{
				year: 2000,
				day:  26,
			},
			want: false,
		},
		"invalid year": {
			args: args{
				year: 2000,
				day:  1,
			},
			want: false,
		},
		"invalid day": {
			args: args{
				year: 2020,
				day:  26,
			},
			want: false,
		},
		"valid year and day": {
			args: args{
				year: 2020,
				day:  1,
			},
			want: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := puzzleExists(test.args.year, test.args.day)
			assert.Equal(t, test.want, got)
		})
	}
}
