package main

import (
	"fmt"
	session1 "hacktiv8/golang-basic/session-1"
	session2 "hacktiv8/golang-basic/session-2"
	"os"
	"strconv"
)

func main() {
	fmt.Println("==== Assignment Looping and Condition ====")
	session1.LoopingAndCondition(10)

	fmt.Println("\n==== Assignment Slice and Looping ====")
	session1.SliceLooping()

	fmt.Println("\n==== Closure Pointer Struct ====")
	session2.ClosurePointerStruct()

	fmt.Println("\n==== Assignment Pertemuan 2 ====")
	// remove the delimeter from the string
	// fmt.Print("\nMasukkan Index Person : ")
	// reader := bufio.NewReader(os.Stdin)
	// // ReadString will block until the delimiter is entered
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("An error occured while reading input. Please try again", err)
	// 	return
	// }

	// input = strings.TrimSuffix(input, "\n")
	args, _ := strconv.Atoi(os.Args[1])
	session2.GetPerson(args)

	// fmt.Println("\nTest Golang")
	// fmt.Println("Hello World! Let's Go!")
}
