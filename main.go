package main

import (
	"fmt"
	session1 "hacktiv8/golang-basic/session-1"
	session2 "hacktiv8/golang-basic/session-2"
	IUserService "hacktiv8/golang-basic/session-3"
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
	if len(os.Args) > 1 {
		args, _ := strconv.Atoi(os.Args[1])
		session2.GetPerson(args)
	}

	fmt.Println("\n==== Assignment Pertemuan 3 ====")
	ourUser := []*IUserService.User{}
	userService := IUserService.NewUserService(ourUser)
	names := []string{"Giva", "Fahmi", "Yusuf"}

	for _, v := range names {
		userName := userService.Register(&IUserService.User{Name: v})
		fmt.Println(userName)
	}

	getUsers := userService.GetUser()
	for _, v := range getUsers {
		fmt.Println("User : ", v.Name)
	}

	// fmt.Println("\nTest Golang")
	// fmt.Println("Hello World! Let's Go!")
}
