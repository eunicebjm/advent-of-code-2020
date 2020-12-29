package main

import "testing"


func TestMultiplyFind2020(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{values: []int{1721,979,366,299,675,1456}},
			//514579,
			241861950,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyFind2020(tt.args.values); got != tt.want {
				t.Errorf("Multiplyfind2020() = %v, want %v", got, tt.want)
			}
		})
	}
}