package queue

type element interface{}

type IQueue interface {
	Push(x element)
	Pop() element
	Size() int
}

type node struct {
	data *element
	next *node
}

type Queue struct {
	head *node
}

func NewQueue() *Queue {
	q := Queue{}
	return &q
}

func (q *Queue) isHeadEmpty() bool {
	return q.head == nil
}

func (q *Queue) getTail() *node {
	var n = q.head
	for n = q.head; n.next != nil; n = n.next {
		// loop
	}

	return n
}

func (q *Queue) setHead(x element) {
	newNode := node{
		data: &x,
		next: nil,
	}

	q.head = &newNode
}

func addAfterTail(tail *node, x element) {
	newNode := node{
		data: &x,
		next: nil,
	}

	tail.next = &newNode
}

func (q *Queue) Push(x element) {
	if q.isHeadEmpty() {
		q.setHead(x)
		return
	}

	// find end of queue
	tail := q.getTail()
	addAfterTail(tail, x)

}

func (q *Queue) Pop() element {
	if q.isHeadEmpty() {
		return nil
	}

	data := q.head.data
	next := q.head.next

	q.head = next

	return *data
}

func (q *Queue) Size() int {
	if q.isHeadEmpty() {
		return 0
	}

	var size = 1
	for n := q.head; n.next != nil; n = n.next {
		size++
	}

	return size
}
