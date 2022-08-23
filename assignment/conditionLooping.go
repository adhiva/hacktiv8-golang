package assignment

import "fmt"

func LoopingAndCondition(maxNumber int) {
	for i := 1; i <= maxNumber; i++ {
		label := "Genap"
		if i%2 == 1 {
			label = "Ganjil"
		}

		fmt.Println(i, label)
	}
}
