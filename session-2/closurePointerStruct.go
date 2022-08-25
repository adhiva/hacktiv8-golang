package session2

import "fmt"

type Person struct {
	Name    string
	Address string
	Job     string
	Reason  string
}

func ClosurePointerStruct() {
	// classMates := []string{
	// 	"Ricky Pieter Palembangan",
	// 	"Thalia Indah Milagrosa",
	// 	"Yosef Brahmantyo Adi Kristanto",
	// 	"Frizky Afif Gautama",
	// 	"Sigit setiawan",
	// 	"Giva",
	// 	"Aulia Nurhady",
	// 	"Fahmi Tajuddin",
	// 	"M. Irfan Maulana",
	// 	"Yusuf Farhan Hasbullah",
	// }
	ourTeamMates := func(persons []*Person) {
		for _, person := range persons {
			fmt.Println("Name : ", person.Name)
		}
	}

	ourTeamMates([]*Person{
		{Name: "Ricky Pieter Palembangan"},
		{Name: "Thalia Indah Milagrosa"},
		{Name: "Yosef Brahmantyo Adi Kristanto"},
		{Name: "Frizky Afif Gautama"},
		{Name: "Sigit setiawan"},
		{Name: "Giva"},
		{Name: "Aulia Nurhady"},
		{Name: "Fahmi Tajuddin"},
		{Name: "M. Irfan Maulana"},
		{Name: "Yusuf Farhan Hasbullah"},
	})
}
