package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGrid_IterateGrid(t *testing.T) {
	lines, err := ReadLines("../test-input.txt")
	require.NoError(t, err)

	type args struct {
		traverseRight int
		traverseDown  int
	}
	tests := []struct {
		name   string
		args   args
		want   int
	}{
		{
			"should return 2 when traverse 1x1",
			args{
				traverseRight: 1,
				traverseDown:  1,
			},
			2,
		},
		{
			"should return 7 when traverse 3x1",
			args{
				traverseRight: 3,
				traverseDown:  1,
			},
			7,
		},
		{
			"should return 3 when traverse 5x1",
			args{
				traverseRight: 5,
				traverseDown:  1,
			},
			3,
		},
		{
			"should return 4 when traverse 7x1",
			args{
				traverseRight: 7,
				traverseDown:  1,
			},
			4,
		},
		{
			"should return 2 when traverse 1x2",
			args{
				traverseRight: 1,
				traverseDown:  2,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Grid{
				lines:  lines,
				width:  11,
				length: 11,
				point:  Point{0,0},
			}
			if got := g.IterateGrid(tt.args.traverseRight, tt.args.traverseDown); got != tt.want {
				t.Errorf("IterateGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}