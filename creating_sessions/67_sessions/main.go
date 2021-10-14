package main

func main() {

}

// ListNode is a struct that holds a value of any type and a poiting to the next node
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func addTwoHugeNumbers(a *ListNode, b *ListNode) *ListNode {
	// the one's place is the only index which
	// can have 3 leading 0s
	// leading 0s not in the one's place are not evaluated
	var x int
	var y int
	final := &ListNode{}

	for {
		if a.Next == nil || b.Next == nil {
			break
		}
		x += a.Value.(int)
		a = a.Next
		y += b.Value.(int)
		b = b.Next
	}

	return final
}

// all indices are a number with 4 digits
// this could be leading 0s
// eg [4] actuall represents 0000
// conversely the number could have trailing zeros
// it is entirely dependent on location of the num
// A & B have both indices being adding together with
// 4 digits before the number is combined for a total
