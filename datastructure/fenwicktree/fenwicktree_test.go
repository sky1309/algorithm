package fenwicktree

import (
	"fmt"
	"testing"
)

func TestFenWickTree(t *testing.T) {
	f := NewFenWickTree([]int{1, 2, 3, 4, 5})
	fmt.Printf("f.Query(4): %v\n", f.Query(4))
	fmt.Printf("f.PreSum(2, 4): %v\n", f.PreSum(2, 4))
}
