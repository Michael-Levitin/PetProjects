// Add two numbers (as linked list)
//https://leetcode.com/problems/add-two-numbers/
//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	numberA := randLinkedNumber(5)
	numberB := randLinkedNumber(7)
	printLinkedNumber(numberA)
	printLinkedNumber(numberB)
	sum := addTwoNumbers(numberA, numberB)
	printLinkedNumber(sum)
}

func randLinkedNumber(n int) *ListNode {
	if n == 0 { // number of digits
		return nil
	}
	rand.Seed(time.Now().UnixNano() + rand.Int63())
	digit := rand.Intn(10)
	//	fmt.Println(digit," n - ", n)
	return &ListNode{digit, randLinkedNumber(n - 1)}
}

func printLinkedNumber(num *ListNode) {
	if num.Next == nil {
		fmt.Println(num.Val)
		return
	}
	fmt.Print(num.Val)
	printLinkedNumber(num.Next)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummyHead = new(ListNode)
	p := l1
	q := l2
	var x, y, overflow int
	curr := dummyHead
	for p != nil || q != nil {
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}

		if q != nil {
			y = q.Val
		} else {
			y = 0
		}

		sum := overflow + x + y
		overflow = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if overflow > 0 {
		curr.Next = &ListNode{overflow, nil}
	}
	return dummyHead.Next
}
