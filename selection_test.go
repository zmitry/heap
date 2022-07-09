package heap

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		list       []int
		left       int
		right      int
		pivotIndex int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{list: []int{5, 3, 6, 4, 7}, left: 0, right: 4, pivotIndex: 2}, want: 3},
		{name: "2", args: args{list: []int{5, 3, 6, 4, 7}, left: 0, right: 4, pivotIndex: 1}, want: 0},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.list, tt.args.left, tt.args.right, tt.args.pivotIndex); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
			fmt.Println(tt.args.list)
		})
	}
}

func TestQuickSelect(t *testing.T) {
	type args struct {
		list  []int
		left  int
		right int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{list: []int{5, 3, 6, 4, 7}, left: 0, right: 4, k: 3}, want: 6},
		{name: "2", args: args{list: []int{5, 4, 3, 2, 1, 4, 8, 2}, left: 0, right: 7, k: 1}, want: 2},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSelect(tt.args.list, tt.args.left, tt.args.right, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}
