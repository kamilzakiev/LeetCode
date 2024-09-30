// https://leetcode.com/problems/all-oone-data-structure/
type AllOne struct {
	DLL  DoublyLinkedList
	MapN map[string]*Node
}

func Constructor() AllOne {
	result := AllOne{}
	result.DLL = DoublyLinkedList{}
	result.MapN = map[string]*Node{}
	return result
}

func (this *AllOne) Inc(key string) {
	if oldNode, exists := this.MapN[key]; !exists {
		// new string
		if this.DLL.head == nil || this.DLL.head.value != 1 {
			this.DLL.InsertAtFront(1)
		}
		this.MapN[key] = this.DLL.head
	} else {
		// existing string
		newCount := oldNode.value + 1 // number of occurence of string related to this node
		oldNode.StringCount--         // number of string with such occurence
        delete(oldNode.Strings, key)
		nextNode := oldNode.next
		prevNode := oldNode.prev
		if oldNode.StringCount == 0 { // need to delete old node
			this.DLL.DeleteNode(oldNode)
		} else {
			prevNode = oldNode
		}
		if nextNode == nil || nextNode.value != newCount { // no node with new value
			if prevNode == nil {
				this.DLL.InsertAtFront(newCount)
				this.MapN[key] = this.DLL.head
			} else {
				this.DLL.InsertAfter(prevNode, newCount)
				this.MapN[key] = prevNode.next
			}
		} else {
			this.MapN[key] = nextNode
		}
	}
	this.MapN[key].StringCount++
    this.MapN[key].Strings[key] = true
}

func (this *AllOne) Dec(key string) {
	oldNode := this.MapN[key]
	// existing string
	newCount := oldNode.value - 1 // number of occurence of string related to this node
	oldNode.StringCount--         // number of string with such occurence
    delete(oldNode.Strings, key)
	prevNode := oldNode.prev
	if oldNode.StringCount == 0 { // need to delete old node
		this.DLL.DeleteNode(oldNode)
	}
    if newCount > 0 {
        if prevNode == nil || prevNode.value != newCount { // no node with new value
            if prevNode == nil {
                this.DLL.InsertAtFront(newCount)
                this.MapN[key] = this.DLL.head
            } else {
                this.DLL.InsertAfter(prevNode, newCount)
                this.MapN[key] = prevNode.next
            }
        } else {
            this.MapN[key] = prevNode
        }
        this.MapN[key].StringCount++
        this.MapN[key].Strings[key] = true
    } else {
        delete(this.MapN, key)
    }
}

func (this *AllOne) GetMaxKey() string {
    //fmt.Println("GetMaxKey", this.DLL.head, this.DLL.tail)
    if this.DLL.tail == nil {
        return ""
    }
    for k := range this.DLL.tail.Strings {
        return k
    }
    return ""
}

func (this *AllOne) GetMinKey() string {
    //fmt.Println("GetMinKey", this.DLL.head, this.DLL.tail)
    if this.DLL.head == nil {
        return ""
    }
    for k := range this.DLL.head.Strings {
        return k
    }
    return ""
}

type Node struct {
	value       int
	prev        *Node
	next        *Node
	StringCount int
    Strings map[string]bool
}

// DoublyLinkedList represents the doubly linked list structure
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// InsertAtFront inserts a new node with a given value at the front of the list
func (dll *DoublyLinkedList) InsertAtFront(value int) {
	newNode := &Node{value: value, Strings: map[string]bool{}}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}
}

// InsertAfter inserts a new node with a given value after the given node
func (dll *DoublyLinkedList) InsertAfter(node *Node, value int) {
	if node == nil {
		fmt.Println("The given node cannot be nil")
		return
	}

	newNode := &Node{value: value, Strings: map[string]bool{}}
	newNode.prev = node
	newNode.next = node.next

	if node.next != nil {
		node.next.prev = newNode
	} else {
		// If node is the tail, we update the tail reference
		dll.tail = newNode
	}

	node.next = newNode
}

// DeleteNode deletes the specified node from the list
func (dll *DoublyLinkedList) DeleteNode(node *Node) {
	if node == nil {
		fmt.Println("The given node cannot be nil")
		return
	}

	// If the node is the head, update the head
	if node == dll.head {
		dll.head = node.next
	}

	// If the node is the tail, update the tail
	if node == dll.tail {
		dll.tail = node.prev
	}

	// Update the next node's prev pointer if not null
	if node.next != nil {
		node.next.prev = node.prev
	}

	// Update the prev node's next pointer if not null
	if node.prev != nil {
		node.prev.next = node.next
	}

	// Clear the pointers in the node (optional, but good practice)
	node.prev = nil
	node.next = nil
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
