package heap

// Heap sss
type Heap struct {
	elements []Comparable
}

// Comparable Element
type Comparable interface {
	GreaterThan(*Comparable) bool
}

func (_ *Heap) findParent(pos int) int {
	if pos <= 0 {
		return -1
	}
	if pos%2 == 0 {
		return pos/2 - 1
	} else {
		return (pos - 1)/2
	}
}

// Add function
func (heap *Heap) Add(element Comparable) {
	heap.elements = append(heap.elements, element)
	pos := len(heap.elements) - 1

	for {
		parent := heap.findParent(pos)
		if parent < 0 {
			return
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
// objects are passed value types
func (heap *Heap) Remove() Comparable{
	length := len(heap.elements)
	if length < 1 {
		return nil
	}
	h := heap.elements[0]
	parent := heap.findParent(length-1)
	if parent < 0 {
		heap.elements = nil
		return h
	}
	heap.elements[0] = heap.elements[length-1]
	heap.elements = heap.elements[:length-1]
	length = length-1
	pos := 0
	for {
		var leftChild Comparable = nil
		var rightChild Comparable = nil
		if 2*pos+1 < length {
			leftChild = heap.elements[2*pos+1]
		} 	
		if 2*pos+2 < length {
			rightChild = heap.elements[2*pos+2]
		}
		var max int
		if leftChild != nil {
			if rightChild != nil && rightChild.GreaterThan(&leftChild){
				max = 2*pos+2
			} else {
				max = 2*pos+1
			}
		} else if rightChild != nil{
			max = 2*pos+2
		} else {
			return h
		}
		maxChild := heap.elements[max]
		if !maxChild.GreaterThan(&heap.elements[pos]) {
			return h
		}
		heap.elements[max] = heap.elements[pos]
		heap.elements[pos] = maxChild
		pos = max
	}
}
