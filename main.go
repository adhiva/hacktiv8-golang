package main

import (
	"fmt"
	session1 "hacktiv8/golang-basic/session-1"
	session2 "hacktiv8/golang-basic/session-2"
)

func main() {
	fmt.Println("==== Assignment Looping and Condition ====")
	session1.LoopingAndCondition(10)

	fmt.Println("\n==== Assignment Slice and Looping ====")
	session1.SliceLooping()

	fmt.Println("\n==== Closure Pointer Struct ====")
	session2.ClosurePointerStruct()

	fmt.Println("\nTest Golang")
	fmt.Println("Hello World! Let's Go!")
}
