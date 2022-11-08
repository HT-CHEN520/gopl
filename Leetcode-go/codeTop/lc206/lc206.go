package main

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func main() {
	var res []int
	l1 := BuildListWithNoHead([]int{1, 2, 3, 4, 5})
	l2 := reverseList(l1)
	for l2 != nil {
		res = append(res, l2.Val)
		l2 = l2.Next
	}
	fmt.Print(res)
}

func reverseList(head *LinkNode) *LinkNode {
	var prev *LinkNode
	p := head
	for p != nil {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return prev
}

func BuildListWithNoHead(nums []int) *LinkNode {
	if len(nums) < 1 {
		return nil
	}

	head := &LinkNode{
		Val:  nums[0],
		Next: nil,
	}
	for i := 1; i < len(nums); i++ {
		insertToTail(head, nums[i])
	}
	return head
}

func insertToTail(head *LinkNode, val int) *LinkNode {
	var cur *LinkNode = head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &LinkNode{
		Val:  val,
		Next: nil,
	}
	return head
}
