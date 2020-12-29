package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetWidth(t *testing.T) {
	t.Run("should get width of text file", func(t *testing.T) {
		lines, err := ReadLines("test-input.txt")
		require.NoError(t, err)
		assert.Equal(t, 11, len(lines[0]))
	})
}

func TestGetLength(t *testing.T) {
	t.Run("should get length of text file", func(t *testing.T) {
		lines, err := ReadLines("test-input.txt")
		require.NoError(t, err)
		assert.Equal(t, 11, len(lines))
	})
}


func TestGrid_IdentifyObject(t *testing.T) {
	lines, err := ReadLines("test-input.txt")
	require.NoError(t, err)

	type fields struct {
		x int
		y int
	}
	type args struct {
		lines []string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"should get tree",
			fields{
				x: 3,
				y: 0,
			},
			args{lines: lines},
			"#",
		},
		{
			"should get open square",
			fields{
				x: 3,
				y: 1,
			},
			args{lines: lines},
			".",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := p.IdentifyObject(tt.args.lines); got != tt.want {
				t.Errorf("IdentifyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIterateGrid(t *testing.T) {
	t.Run("should return tree count as 7", func(t *testing.T) {
		lines, err := ReadLines("test-input.txt")
		require.NoError(t, err)
		treeCount := IterateGrid(lines)
		assert.Equal(t, 7, treeCount)
	})
}