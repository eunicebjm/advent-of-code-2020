package main

import (
	"testing"
)

func TestGetValidPasswordCount(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{[]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValidPasswordCount(tt.args.lines); got != tt.want {
				t.Errorf("GetValidPasswordCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_record_IsValid(t *testing.T) {
	type fields struct {
		rule     string
		password string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"case1",
			fields{"1-3 a", "abcde"},
			true,
		},
		{
			"case2",
			fields{"1-3 b", "cdefg"},
			false,
		},
		{
			"case3",
			fields{"2-9 c", "ccccccccc"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := record{
				rule:     tt.fields.rule,
				password: tt.fields.password,
			}
			if got := r.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}