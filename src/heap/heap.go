package heap

// Heap sss
type Heap struct {
	elements []Comparable
}

// Comparable Element
type Comparable interface {
	GreaterThan(*Comparable) bool
}

// Add function
func (heap *Heap) Add(element Comparable) {
	heap.elements = append(heap.elements, element)
	pos := len(heap.elements) - 1

	for {
		if pos == 0 {
			return
		}
		var parent int
		if pos%2 == 0 { // right child
			parent = (pos - 2) / 2
		} else { // left child
			parent = (pos - 1) / 2
		}
		newEle := heap.elements[pos]
		if !newEle.GreaterThan(&heap.elements[parent]) {
			break
		}
		heap.elements[pos] = heap.elements[parent]
		heap.elements[parent] = newEle
		pos = parent
	}
}
