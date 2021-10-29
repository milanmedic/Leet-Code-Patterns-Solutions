package queue

type Element struct {
	value interface{}
	next  *Element
}

type Queue struct {
	head        *Element
	end         *Element
	lookupTable map[interface{}]bool
	Count       int
}

func NewQueue() *Queue {
	return &Queue{lookupTable: map[interface{}]bool{}}
}

func (q *Queue) Enqueue(value interface{}) {
	q.Count++
	if q.head == nil {
		q.head = &Element{value: value}
		q.lookupTable[value] = true
		q.end = q.head
		return
	}

	q.end.next = &Element{value: value}
	q.lookupTable[value] = true
	q.end = q.end.next
}

func (q *Queue) Dequeue() interface{} {
	dequeuedElem := q.head
	q.head = q.head.next
	delete(q.lookupTable, dequeuedElem.value)
	q.Count--
	return dequeuedElem.value
}

func (q *Queue) Contains(value interface{}) bool {
	_, found := q.lookupTable[value]
	return found
}
