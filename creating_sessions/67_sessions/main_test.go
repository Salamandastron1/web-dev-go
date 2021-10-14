package main

import (
	"fmt"
	"testing"
)

type test struct {
	data data
	want *ListNode
}
type data struct {
	a *ListNode
	b *ListNode
}

var testCases []test = []test{
	{
		data: data{
			a: populateLinkedList([]int{9876, 5432, 1999}),
			b: populateLinkedList([]int{1, 8001}),
		},
		want: populateLinkedList([]int{9876, 5434, 0}),
	},
	{
		data: data{
			a: populateLinkedList([]int{123, 4, 5}),
			b: populateLinkedList([]int{100, 100, 100}),
		},
		want: populateLinkedList([]int{223, 104, 105}),
	},
	{
		data: data{
			a: populateLinkedList([]int{0}),
			b: populateLinkedList([]int{0}),
		},
		want: populateLinkedList([]int{0}),
	},
	{
		data: data{
			a: populateLinkedList([]int{1234, 1234, 0}),
			b: populateLinkedList([]int{0}),
		},
		want: populateLinkedList([]int{1234, 1234, 0}),
	},
	{
		data: data{
			a: populateLinkedList([]int{0}),
			b: populateLinkedList([]int{1234, 1234, 0}),
		},
		want: populateLinkedList([]int{1234, 1234, 0}),
	},
	{
		data: data{
			a: populateLinkedList([]int{1}),
			b: populateLinkedList([]int{9998, 9999, 9999, 9999, 9999, 9999}),
		},
		want: populateLinkedList([]int{9999, 0, 0, 0, 0, 0}),
	},
	{
		data: data{
			a: populateLinkedList([]int{1}),
			b: populateLinkedList([]int{9999, 9999, 9999, 9999, 9999, 9999}),
		},
		want: populateLinkedList([]int{1, 0, 0, 0, 0, 0, 0}),
	},
	{
		data: data{
			a: populateLinkedList([]int{8339, 4510}),
			b: populateLinkedList([]int{2309}),
		},
		want: populateLinkedList([]int{8339, 6819}),
	},
}

func TestAddTwoHugeNums(t *testing.T) {
	for i, v := range testCases {
		got := addTwoHugeNumbers(v.data.a, v.data.b)
		for {
			fmt.Println("Test #", i)
			fmt.Println("Got:", got)
			fmt.Println("Want:", v.want)
			if got.Next != v.want.Next {
				t.Errorf("Test #%v\nGot NOT equal to want\nGot:%v\nWant:%v", i, got.Value, v.want.Value)
				break
			}
			if got.Value != v.want.Value {
				t.Errorf("Test #%v\nGot NOT equal to want\nGot:%v\nWant:%v", i, got.Value, v.want.Value)
				break
			}

			if got.Next == nil && v.want.Next == nil {
				break
			}
			got = got.Next
			v.want = v.want.Next
		}
	}
}

func populateLinkedList(xi []int) *ListNode {
	head := &ListNode{}
	current := head
	for i, v := range xi {
		current.Value = v
		if i != len(xi)-1 {
			current.Next = &ListNode{}
			current = current.Next
		}
	}

	return head
}
