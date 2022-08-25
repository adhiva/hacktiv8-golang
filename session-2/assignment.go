package session2

import "fmt"

func GetPerson(index int) {
	names := []*Person{
		{Name: "Ricky Pieter Palembangan", Address: "Jakarta", Job: "Programmer", Reason: "Growth Together"},
		{Name: "Thalia Indah Milagrosa", Address: "Jakarta", Job: "Programmer", Reason: "Growth Together"},
		{Name: "Yosef Brahmantyo Adi Kristanto", Address: "Jakarta", Job: "Programmer", Reason: "Growth Together"},
		{Name: "Frizky Afif Gautama", Address: "Jakarta", Job: "Programmer", Reason: "Growth Together"},
		{Name: "Sigit setiawan", Address: "Jakarta", Job: "Programmer", Reason: "Growth Together"},
		{Name: "Giva", Address: "Jakarta", Job: "Programmer", Reason: "Growth Together"},
		{Name: "Aulia Nurhady", Address: "Jakarta", Job: "Programmer", Reason: "Growth Together"},
		{Name: "Fahmi Tajuddin", Address: "Jakarta", Job: "Programmer", Reason: "Growth Together"},
		{Name: "M. Irfan Maulana", Address: "Jakarta", Job: "Programmer", Reason: "Growth Together"},
		{Name: "Yusuf Farhan Hasbullah", Address: "Jakarta", Job: "Programmer", Reason: "Growth Together"},
	}

	if index > len(names) || index < 0 {
		fmt.Print("ERROR Undifined Index!")
		return
	}

	fmt.Println(
		"Nama : ", names[index].Name,
		"\nAlamat : ", names[index].Address,
		"\nPekerjaan : ", names[index].Job,
		"\nAlasan ikut kelas Golang : ", names[index].Reason,
	)
	return
}
