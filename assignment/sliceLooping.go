package assignment

import "fmt"

func SliceLooping() {
	names := []string{
		"Ricky Pieter Palembangan",
		"Thalia Indah Milagrosa",
		"Yosef Brahmantyo Adi Kristanto",
		"Frizky Afif Gautama",
		"Sigit setiawan",
		"Giva",
		"Aulia Nurhady",
		"Fahmi Tajuddin",
		"M. Irfan Maulana",
		"Yusuf Farhan Hasbullah",
	}

	for _, name := range names {
		fmt.Println("Name : ", name)
	}
}
