package common

type Queue struct {
	elements []interface{}
}

func NewQueue() Queue {
	return Queue{
		elements: make([]interface{}, 0),
	}
}

func (q Queue) Push(i interface{}) {
	if i == nil {
		panic("can't push nil element")
	}
	q.elements = append(q.elements, i)
}

func (q Queue) Pop() interface{} {
	if len(q.elements) == 0 {
		return nil
	}
	item := q.elements[0]
	q.elements = q.elements[1:]
	return item
}

// 优先级队列
type PriorityQueue struct {
	data []priorityNode
}

type priorityNode struct {
	priority int
	value    interface{}
}

func NewPriorityQueue() PriorityQueue {
	return PriorityQueue{
		data: make([]priorityNode, 0),
	}
}

func (q *PriorityQueue) Push(priority int, val interface{}) {
	node := priorityNode{
		priority: priority,
		value:    val,
	}
	q.data = append(q.data, node)
	for i := len(q.data) - 2; i >= 0; i-- {
		if q.data[i].priority < q.data[i+1].priority {
			tmp := q.data[i]
			q.data[i] = q.data[i+1]
			q.data[i+1] = tmp
		}
	}
}

func (q *PriorityQueue) Pop() interface{} {
	if len(q.data) > 0 {
		ret := q.data[0]
		q.data = q.data[1:]
		return ret.value
	} else {
		panic("queue is empty")
	}
}
