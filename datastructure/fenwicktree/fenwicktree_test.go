package fenwicktree

import (
	"fmt"
	"testing"
)

func TestFenWickTree(t *testing.T) {
	f := NewFenWickTree([]int{1, 2, 3, 4, 5})
	fmt.Printf("f.PreSum(4): %v\n", f.PreSum(2))
	fmt.Printf("f.Range(2, 4): %v\n", f.Range(2, 3))
}

func ExampleFenWickTree() {
	fmt.Println("===")
	// output:
}
