package circlequeue

import (
	"fmt"
	"testing"
)

func TestCircleQueue(t *testing.T) {
	cq := NewCircleQueue[int](10)
	for i := 0; i < 14; i++ {
		cq.Enqueue(i)
		fmt.Println(cq)
	}

	for i := 0; i < 5; i++ {
		val, ok := cq.Dequeue()
		fmt.Printf("dequeu val=%+v, ok=%v\n", val, ok)
	}

	fmt.Println(cq)
}
