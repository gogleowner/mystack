package main

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

func main() {
	//example_1_list()

	//example_2_heap()
	//Example_priorityQueue()

	exmaple_3_ring()
}

// 10 -> 20 -> 30
func example_1_list() {
	l := list.New()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	fmt.Println("Front ", l.Front().Value)
	fmt.Println("Back ", l.Back().Value)

	for element := l.Front(); element != nil; element = element.Next() {
		fmt.Printf("%d -> ", element.Value)
	}

	fmt.Println("\nreverse order")
	for element := l.Back(); element != nil; element = element.Prev() {
		fmt.Printf("%d <- ", element.Value)
	}
}



type MinHeap []int // heap을 int slice로 정의

func (h *MinHeap) Push(element interface{}) {
	fmt.Println("Push", element)
	*h = append(*h, element.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	element := old[n-1]
	*h = old[0 : n-1]

	return element
}


func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(num1, num2 int) bool { // 대소관계 판단
	r := h[num1] < h[num2]
	fmt.Printf("less %d < %d %t\n", num1, num2, r)
	return r
}

func (h MinHeap) Swap(num1, num2 int) {
	fmt.Printf("Swap %d %d\n", h[num1], h[num2])
	h[num1], h[num2] = h[num2], h[num1]
}


func example_2_heap() {
	data := new(MinHeap)
	heap.Init(data)

	heap.Push(data, 5)
	heap.Push(data, 2)
	heap.Push(data, 7)
	heap.Push(data, 3)

	fmt.Println(data, "최소값 :", (*data)[0])
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func Example_priorityQueue() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}

	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
	// Output:
	// 05:orange 04:pear 03:banana 02:apple
}




func exmaple_3_ring() {
	data := []string{"Maria", "John", "Andrew", "James"}

	r := ring.New(len(data))
	for i := 0; i < r.Len(); i++ {
		r.Value = data[i]
		r = r.Next()
	}

	r.Do(func (x interface{}) {
		fmt.Println(x)
	})

	fmt.Println("Move forward : 1")
	r = r.Move(1) // 시계방향
	fmt.Println("Curr :", r.Value)
	fmt.Println("Prev :", r.Prev().Value)
	fmt.Println("Next :", r.Next().Value)

	fmt.Println("Move backward : -1")
	r = r.Move(-1) // 반시계방향
	fmt.Println("Curr :", r.Value)
	fmt.Println("Prev :", r.Prev().Value)
	fmt.Println("Next :", r.Next().Value)
}

