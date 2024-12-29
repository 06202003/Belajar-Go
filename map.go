package main
import "fmt"

func main(){
	person := map[string]string{
		"name":"Yehezkiel",
		"nrp":"2172003",
	}
	fmt.Println(person)

	var biodata map[string]string = map[string]string{}

	biodata["name"]="David"
	biodata["address"]="Cimahi"
	biodata["gender"]="Male"

	fmt.Println(biodata)
	fmt.Println(len(biodata))
	bioName := biodata["name"]
	fmt.Println(bioName)

	// MAKE

	book := make(map[string]string)
	book["title"] = "Harry Potter"
	book["author"] = "JK Rowling"
	book["publisher"] = "Gramedia"

	fmt.Println(book)
	delete(book,"publisher")
	fmt.Println(book)


}