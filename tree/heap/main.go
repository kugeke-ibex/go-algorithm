package main

import (
	"container/heap"
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

	fmt.Println("--------------------------------")
	fmt.Println("container/heapを使用した最大ヒープ")

	// MaxHeap型のヒープを作成
	maxHeap := MaxHeap{}
	// heap.Interfaceを初期化
	heap.Init(&maxHeap)

	heap.Push(&maxHeap, 3)
	heap.Push(&maxHeap, 1)
	heap.Push(&maxHeap, 4)
	heap.Push(&maxHeap, 2)
	heap.Push(&maxHeap, 4)
	heap.Push(&maxHeap, 6)
	heap.Push(&maxHeap, 10)
	heap.Push(&maxHeap, 6)
	heap.Push(&maxHeap, 9)
	heap.Push(&maxHeap, 6)
	fmt.Println(maxHeap)
	fmt.Println("ヒープのサイズ:", maxHeap.Len()) // 10

	// 要素を取り出す（最大値から順に取り出される）
	for maxHeap.Len() > 0 {
		// heap.Pop()はPopメソッドを呼び出し、最後にheap.Fixで調整します。
		fmt.Printf("%d ", heap.Pop(&maxHeap)) // 10 9 8 7 6 5 4 3 2 1
	}
	fmt.Println()

	words := []string{
		"python", "c", "java", "go", "python", "c", "go", "python",
	}

	fmt.Println("--------------------------------")
	fmt.Println("文字列の出現回数トップNをヒープで取得")

	// 出現回数をカウント
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	// ヒープを作成
	wordHeap := WordCountHeap{}
	heap.Init(&wordHeap)

	// カウント結果をヒープに追加
	for word, count := range wordCount {
		heap.Push(&wordHeap, WordCount{Word: word, Count: count})
	}

	// トップNを取得（例: トップ3）
	N := 3
	fmt.Printf("トップ%dの文字列:\n", N)
	topN := getTopN(&wordHeap, N)
	for i, wc := range topN {
		fmt.Printf("%d位: %s (出現回数: %d)\n", i+1, wc.Word, wc.Count)
	}

}

type MiniHeap struct {
	heap        []int
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

// IntHeap はintのスライスで、heap.Interfaceを実装します。
type MaxHeap []int

// Len: ヒープの要素数を返します。
func (h MaxHeap) Len() int { return len(h) }

// Less: 最大ヒープの場合、h[i] > h[j]ならtrueを返します。
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap: 要素を交換します。
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push: ヒープに要素を追加します（スライスにappendするだけ）。
func (h *MaxHeap) Push(x any) {
	// x.(int)で型アサーションしてintとして扱います。
	*h = append(*h, x.(int))
}

// Pop: ヒープから最小要素を取り除き返します（スライスからpopする）。
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]     // 最後の要素を取得
	*h = old[0 : n-1] // スライスを短くする
	return x
}

// WordCount は文字列とその出現回数を保持する構造体
type WordCount struct {
	Word  string
	Count int
}

// WordCountHeap はWordCountのスライスで、heap.Interfaceを実装します（最大ヒープ）
type WordCountHeap []WordCount

// Len: ヒープの要素数を返します
func (h WordCountHeap) Len() int { return len(h) }

// Less: 最大ヒープの場合、出現回数が多い順に並べます
// 出現回数が同じ場合は、文字列の辞書順で比較します
func (h WordCountHeap) Less(i, j int) bool {
	if h[i].Count == h[j].Count {
		return h[i].Word < h[j].Word // 同じ出現回数の場合は辞書順
	}
	return h[i].Count > h[j].Count // 出現回数の多い順
}

// Swap: 要素を交換します
func (h WordCountHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push: ヒープに要素を追加します
func (h *WordCountHeap) Push(x any) {
	*h = append(*h, x.(WordCount))
}

// Pop: ヒープから最大要素を取り除き返します
func (h *WordCountHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// getTopN はヒープからトップNの要素を取得します
func getTopN(h *WordCountHeap, n int) []WordCount {
	result := make([]WordCount, 0, n)

	for i := 0; i < n && h.Len() > 0; i++ {
		result = append(result, heap.Pop(h).(WordCount))
	}

	return result
}
