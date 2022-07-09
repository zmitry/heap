package heap

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

// [1,_,_]
// [_,2,3,_]
// [_,2,_,4,5]
// [_,_,3,_,_,_,5]
func Parent(k int) int {
	return (k - 1) / 2
}

func Less[N constraints.Ordered](a, b N) bool {
	return a < b
}

func NewHeap[N constraints.Ordered]() *Heap[N] {
	return &Heap[N]{
		Less: Less[N],
	}
}

type Heap[N constraints.Ordered] struct {
	arr  []N
	Less func(a N, b N) bool
}

func (h *Heap[N]) Len() int {
	return len(h.arr)
}

func (h *Heap[N]) Insert(item N) {
	h.arr = append(h.arr, item)
	h.up(h.Len() - 1)
	fmt.Println(h.arr)

}

func (h *Heap[N]) Pop() N {
	if len(h.arr) == 0 {
		var t N
		return t
	}
	l := h.arr[0]
	h.arr = h.arr[1:]
	down(h.arr, 0, h.Len()-1, h.Less)
	return l
}

func (h *Heap[N]) Top(item N) (N, error) {
	if len(h.arr) == 0 {
		var t N
		return t, errors.New("empty heap")
	}

	return h.arr[0], nil
}

func (h *Heap[N]) up(i int) {
	for {
		p := Parent(i)
		if p == i || h.Less(h.arr[p], h.arr[i]) {
			break
		}
		h.arr[p], h.arr[i] = h.arr[i], h.arr[p]
		i = p
	}
}

func HeapSort[N constraints.Ordered](arr []N, More func(a, b N) bool) {
	end := len(arr) - 1

	for start := end/2 - 1; start >= 0; start-- {
		down(arr, start, end, More)
	}

	for end > 0 {
		arr[end], arr[0] = arr[0], arr[end]
		end--
		down(arr, 0, end, More)
	}
}

func down[N constraints.Ordered](arr []N, i int, count int, Less func(a, b N) bool) {
	child := 2*i + 1
	for child <= count && child >= 0 {
		if j2 := child + 1; j2 <= count && Less(arr[j2], arr[child]) {
			child = j2
		}
		if !Less(arr[child], arr[i]) {
			return
		}
		arr[i], arr[child] = arr[child], arr[i]
		i = child
		child = 2*i + 1
	}
}
