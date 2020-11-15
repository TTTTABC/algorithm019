/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	ret := -1
	var h IntHeap
	heap.Init(&h)
	heap.Push(&h, 1)
	for n > 0 {
		for h[0] == ret {
			heap.Pop(&h)
		}
		ret = heap.Pop(&h).(int)
		heap.Push(&h, 2*ret)
		heap.Push(&h, 3*ret)
		heap.Push(&h, 5*ret)
		n--
	}
	return ret
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int){ 
	h[i], h[j] = h[j], h[i] 
}

func (h *IntHeap) Push(x interface{}) { 
	*h = append(*h, x.(int)) 
}

func (h *IntHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

// @lc code=end

