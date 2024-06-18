package fenwicktree

/*
	适用于常更新的区间和
	1  0001 1  [1,1]
	2  0010	2  [1,2]
	3  0011 1  [3,3]
	4  0100 4  [1,4]
	5  0101 1  [5,5]
	6  0110 2  [5,6]
	7  0111 1  [7,7]
	8  1000 8  [1,2,3,4,5,6,7,8]
	9  1001 1  [9,9]
	10 1010 2  [9,10]
	11 1011 1  [11,11]
	12 1100 4  [9,10,11,12]
	13 1101 1  [13,13]
	14 1110 2  [13,14]
	15 1111 1  [15,15]
*/

func NewFenWickTree(arr []int) *FenWickTree {
	n := len(arr)
	tree := &FenWickTree{
		bits: make([]int, n+1),
	}
	for i, val := range arr {
		tree.Add(i, val)
	}
	return tree
}

type FenWickTree struct {
	bits []int
}

func (f *FenWickTree) Add(i int, val int) {
	for k := i + 1; k < len(f.bits); k += (k & -k) {
		f.bits[k] += val
	}
}

func (f *FenWickTree) PreSum(i int) int {
	sum := 0
	for k := i + 1; k > 0; k -= (k & -k) {
		sum += f.bits[k]
	}
	return sum
}

func (f *FenWickTree) Range(l, r int) int {
	return f.PreSum(r) - f.PreSum(l-1)
}
