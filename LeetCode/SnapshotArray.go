//https://leetcode.com/problems/snapshot-array/
/**
Implement a SnapshotArray that supports the following interface:

SnapshotArray(int length) initializes an array-like data structure with the given length.  Initially, each element equals 0.
void set(index, val) sets the element at the given index to be equal to val.
int snap() takes a snapshot of the array and returns the snap_id: the total number of times we called snap() minus 1.
int get(index, snap_id) returns the value at the given index, at the time we took the snapshot with the given snap_id

Your SnapshotArray object will be instantiated and called as such:
obj := Constructor(length);
obj.Set(index,val);
param_2 := obj.Snap();
param_3 := obj.Get(index,snap_id);

Input: ["SnapshotArray","set","snap","set","get"]
[[3],[0,5],[],[0,6],[0,0]]
Output: [null,null,0,null,5]
Explanation:
SnapshotArray snapshotArr = new SnapshotArray(3); // set the length to be 3
snapshotArr.set(0,5);  // Set array[0] = 5
snapshotArr.snap();  // Take a snapshot, return snap_id = 0
snapshotArr.set(0,6);
snapshotArr.get(0,0);  // Get the value of array[0] with snap_id = 0, return 5
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type SnapshotArray struct {
	slice       []int
	sliceLength int
	snapLength  int
	snap        []map[int]int // contains snapshots of slice
}

func Constructor(length int) SnapshotArray {
	snapshotArray := new(SnapshotArray)
	snapshotArray.slice = make([]int, length) // don't really need it
	snapshotArray.sliceLength = length
	//Adding new empty map to the slice
	snapshotArray.snap = append(snapshotArray.snap, make(map[int]int))
	return *snapshotArray
}

func (this *SnapshotArray) Set(index int, val int) {
	this.slice[index] = val
	this.snap[this.snapLength][index] = val
}

func (this *SnapshotArray) Snap() int {
	this.snapLength++
	this.snap = append(this.snap, make(map[int]int))
	return this.snapLength - 1
}

func (this *SnapshotArray) Get(index int, snapId int) int {
	for i := snapId; i >= 0; i-- {
		if val, exist := this.snap[i][index]; exist {
			return val
		}
	}
	return 0
}

func selector(commands []string, data [][]int) {
	length := len(commands)
	if length != len(data) {
		fmt.Println("commands and data lengths don't match")
		time.Sleep(5 * time.Second)
		return
	}
	answers := make([]int, length)
	obj := Constructor(data[0][0])
	for i := 1; i < length; i++ {
		switch commands[i] {
		case "set":
			obj.Set(data[i][0], data[i][1])
		case "snap":
			answers[i] = obj.Snap()
		case "get":
			answers[i] = obj.Get(data[i][0], data[i][1])
		}
	}
	fmt.Println(answers)
}

func decodeJson(output interface{}, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't open the file", err)
	}
	// using file and address of var. var must be same structure as Json
	if err := json.NewDecoder(file).Decode(&output); err == nil {
		fmt.Println("Decode successful")
	} else {
		fmt.Println("Decode unsuccessful", err)
	}
	file.Close()
	//fmt.Println(output)
} // usage decodeJson(&someVar, "filename.ext")

func main() {
	//commands:= []string{"SnapshotArray","set","snap","snap","snap","get","snap","snap","get"}
	//data:= [][]int{{1},{0,15},{},{},{},{0,2},{},{},{0,0}}
	////fmt.Println(data, commands)
	//selector(commands, data)

	var commands []string
	decodeJson(&commands, "SnapshotArray-Commands.json")
	var data [][]int
	decodeJson(&data, "SnapshotArray-Data.json")
	//fmt.Println(len(commands), len(data),  "DONE")
	selector(commands, data)
}
