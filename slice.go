package main
import "fmt"

func main(){
	names:=[...]string{
		"Yehezkiel",
		"David",
		"Setiawan",
		"Joko",
		"Ahmad",
		"Budi",
		"Windah",
		"Basudara",
	}

	slice1:= names[4:6]
	fmt.Println(slice1)

	slice2:= names[:3]
	fmt.Println(slice2)

	slice3:= names[3:]
	fmt.Println(slice3)

	slice4:= names[:]
	fmt.Println(slice4)

	// alternative
	var slice5 []string = names[1:5]
	fmt.Println(slice5)

	// slice operation
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	// append
	days:=[...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}
	
	daySlice1:= days[5:]
	fmt.Println(daySlice1)
	daySlice1[0] = "saturday"
	daySlice1[1] = "sunday"
	fmt.Println(days)

	daySlice2:= append(daySlice1,"Libur Lagi!")
	daySlice2[0] = "sabtu njay"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	// NEW SLICE without ARRAY first
	// make(tipeData,len,cap)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Yehezkiel"
	newSlice[1] = "David"
	

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// tambah data
	newSlice2 := append(newSlice,"Setiawan")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// COPY SLICE

	fromSlice := days[:]
	toSlice := make([]string,len(fromSlice),cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}