package circlequeue

import (
	"fmt"
	"strings"
)

func NewCircleQueue[T comparable](maxSize int) *CircleQueue[T] {
	return &CircleQueue[T]{
		data:       make([]T, maxSize),
		maxSize:    maxSize,
		size:       0,
		enqueueIdx: 0,
		dequeueIdx: 0,
	}
}

type CircleQueue[T comparable] struct {
	data       []T
	maxSize    int // 循环队列的最大数量
	size       int // 当前队列元素的个数
	enqueueIdx int // 入队列位置
	dequeueIdx int // 出队列位置
}

func (cq *CircleQueue[T]) Enqueue(t T) {
	if cq.IsFull() {
		// remove first
		cq.Dequeue()
	}

	cq.size++
	cq.data[cq.enqueueIdx] = t

	cq.enqueueIdx++
	if cq.enqueueIdx >= cq.maxSize {
		cq.enqueueIdx = 0
	}
}

func (cq *CircleQueue[T]) Dequeue() (T, bool) {
	var value T
	if cq.size == 0 {
		return value, false
	}

	cq.size--
	value = cq.data[cq.dequeueIdx]

	cq.dequeueIdx++
	if cq.dequeueIdx == cq.maxSize {
		cq.dequeueIdx = 0
	}
	return value, true
}

func (cq *CircleQueue[T]) IsFull() bool {
	return cq.size >= cq.maxSize
}

func (cq *CircleQueue[T]) Clear() {
	cq.data = make([]T, cq.maxSize)
	cq.size = 0
	cq.enqueueIdx = 0
	cq.dequeueIdx = 0
}

func (cq *CircleQueue[T]) Values() []T {
	values := make([]T, cq.size)
	for i := 0; i < cq.size; i++ {
		values[i] = cq.data[(cq.dequeueIdx+i)%cq.maxSize]
	}
	return values
}

func (cq *CircleQueue[T]) String() string {
	str := "CircleQueue\n"
	var values []string
	for _, value := range cq.Values() {
		values = append(values, fmt.Sprintf("%+v", value))
	}
	str += strings.Join(values, ",")
	return str
}
