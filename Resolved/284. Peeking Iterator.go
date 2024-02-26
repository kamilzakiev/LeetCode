// https://leetcode.com/problems/peeking-iterator
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
    Iter *Iterator
    PeekVal int
    HasNext bool
}

func Constructor(iter *Iterator) *PeekingIterator {
    result := PeekingIterator{Iter: iter}
    if iter.hasNext() {
        result.PeekVal = iter.next()
        result.HasNext = true
    } else {
        result.HasNext = false
    }
    return &result
}

func (this *PeekingIterator) hasNext() bool {
    return this.HasNext
}

func (this *PeekingIterator) next() int {
    if this.Iter.hasNext() {
        newVal := this.Iter.next()
        oldVal := this.PeekVal
        this.PeekVal = newVal
        return oldVal
    } else {
        this.HasNext = false
        return this.PeekVal
    }
}

func (this *PeekingIterator) peek() int {
    return this.PeekVal
}
