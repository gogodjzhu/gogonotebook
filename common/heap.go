package common

import "errors"

/**
小顶堆
*/
type Heap struct {
	data []int
}

func NewHeap(vals ...int) *Heap {
	h := &Heap{}
	for _, v := range vals {
		h.Put(v)
	}
	return h
}

func (h *Heap) Put(val int) {
	h.data = append(h.data, val)
	h.popup(len(h.data) - 1)
}

func (h *Heap) Poll() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("heap is empty")
	}
	v := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return v, nil
}

/**
小顶堆排序后为升序
*/
func (h *Heap) Sort() {
	for i := 0; i < len(h.data); i++ {
		tmp := h.data[0]
		h.data[0] = h.data[len(h.data)-i-1]
		h.data[len(h.data)-i-1] = tmp

		h.sink(0, len(h.data)-i-1)
	}
}

func (h *Heap) ToList() *ListNode {
	v, err := h.Poll()
	if err != nil {
		return nil
	}
	return &ListNode{
		Val:  v,
		Next: h.ToList(),
	}
}

/**
将下标为i的节点下沉, minFixedIdx为已经固定的最小下标(大于此下标的均已固定不需要操作)
*/
func (h *Heap) sink(i, fixedIdx int) {
	if i >= len(h.data)/2 || i*2+1 >= fixedIdx {
		// 叶子节点 or 所有子节点均为fixed
		return
	}
	minChildIdx := i*2 + 1
	if i*2+2 < fixedIdx { // 右节点未固定时才参与比较
		if h.data[i*2+1] < h.data[i*2+2] {
			minChildIdx = i*2 + 1
		} else {
			minChildIdx = i*2 + 2
		}
	}

	if h.data[i] > h.data[minChildIdx] {
		Swap(h.data, i, minChildIdx)
		h.sink(minChildIdx, fixedIdx)
	}
}

func (h *Heap) popup(i int) {
	if i == 0 {
		// 到达根节点
		return
	}
	parentIdx := (i - 1) / 2
	if h.data[parentIdx] > h.data[i] {
		Swap(h.data, i, parentIdx)
		h.popup(parentIdx)
	}
}
