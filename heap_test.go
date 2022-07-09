package heap_test

import (
	"leet-go/heap"
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{name: "1", args: []int{6, 4, 3, 7}, want: []int{3, 4, 6, 7}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := heap.NewHeap[int]()
			for _, v := range tt.args {
				h.Insert(v)
			}
			res := []int{}
			for h.Len() > 0 {
				res = append(res, h.Pop())
			}
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("Heap() = %v, want %v", res, tt.want)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{name: "1", args: []int{5, 3, 6, 4, 7}, want: []int{3, 4, 5, 6, 7}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			heap.HeapSort(tt.args, func(a int, b int) bool {
				return a > b
			})
			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("Heap() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
