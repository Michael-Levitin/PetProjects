// Not really from LeetCode, but related

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
	numberA := randLinkedNumber(10)
	printLinkedNumber(numberA)
	printLinkedNumber(reverseLinkedNumber(numberA))
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

func reverseLinkedNumber(num *ListNode) *ListNode {
	ptrToPrev := num //last digit ...
	num, num.Next = num.Next, nil

	for num.Next != nil {
		num, num.Next, ptrToPrev = num.Next, ptrToPrev, num
	} //swapping all 3 simultaneously
	num.Next = ptrToPrev // ... and first digit need special treatment
	return num
}

func printLinkedNumber(num *ListNode) {
	if num.Next == nil {
		fmt.Println(num.Val)
		return
	}
	fmt.Print(num.Val)
	printLinkedNumber(num.Next)
}
