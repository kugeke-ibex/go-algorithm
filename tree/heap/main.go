package main

import (
	"fmt"
	"math"
)

func main() {
	minHeap := MiniHeap{}
	minHeap.Init()

	minHeap.push(5)
	minHeap.push(6)
	minHeap.push(2)
	minHeap.push(9)
	minHeap.push(13)
	minHeap.push(11)
	minHeap.push(1)
	fmt.Println(minHeap.heap)
	fmt.Println(*minHeap.pop())
	fmt.Println(minHeap.heap)




}

type MiniHeap struct {
	heap []int
	currentSize int
}

func (h *MiniHeap) Init() {
	h.heap = []int{-1 * math.MaxInt64}
	h.currentSize = 0
}

func (h MiniHeap) parentIndex(index int) int {
	return index / 2
}

func (h MiniHeap) leftChildIndex(index int) int {
	return index * 2
}

func (h MiniHeap) rightChildIndex(index int) int {
	return (index * 2) + 1
}

func (h *MiniHeap) swap(index1, index2 int) {
	h.heap[index1], h.heap[index2] = h.heap[index2], h.heap[index1]
}

func (h *MiniHeap) heapifyUp(index int) {
	for h.parentIndex(index) > 0 {
		if h.heap[index] < h.heap[h.parentIndex(index)] {
			h.swap(index, h.parentIndex(index))
		}

		index = h.parentIndex(index)
	}
}

func (h *MiniHeap) push(value int) {
	h.heap = append(h.heap, value)
	h.currentSize++
	h.heapifyUp(h.currentSize)
}

func (h *MiniHeap) minChildIndex(index int) int {
	if h.rightChildIndex(index) > h.currentSize {
		return h.leftChildIndex(index)
	} else {
		if h.heap[h.leftChildIndex(index)] < h.heap[h.rightChildIndex(index)] {
			return h.leftChildIndex(index)
		} else {
			return h.rightChildIndex(index)
		}
	}
}

func (h *MiniHeap) heapifyDown(index int) {
	for h.leftChildIndex(index) <= h.currentSize {
		minChildIndex := h.minChildIndex(index)
		if h.heap[index] > h.heap[minChildIndex] {
			h.swap(index, minChildIndex)
		}

		index = minChildIndex
	}
}

func (h *MiniHeap) pop() *int {
	if h.currentSize < 1 {
		return nil
	}

	root := h.heap[1]
	data := h.heap[h.currentSize]
	h.heap = h.heap[:h.currentSize]
	if len(h.heap) == 1 {
		return &root
	}

	h.heap[1] = data
	h.currentSize--
	h.heapifyDown(1)

	return &root
}
