package main
import "fmt"

func main(){
	var biodata[3] string
	var fullName,nrp,gender string
	var uts,uas int

	fmt.Print("Full name: ")
	fmt.Scanln(&fullName)
	fmt.Print("NRP: ")
	fmt.Scanln(&nrp)
	fmt.Print("gender (male/female): ")
	fmt.Scanln(&gender)

	biodata[0]= fullName
	biodata[1]= nrp
	biodata[2]= gender

	fmt.Println(biodata)
	fmt.Println("---------------------------")
	fmt.Println(biodata[1],biodata[0],biodata[2])

	// Langsung tanpa set
	fmt.Print("Nilai KAT: ")
	fmt.Scan(&kat)
	fmt.Print("Nilai UTS: ")
	fmt.Scan(&uts)
	fmt.Print("Nilai UAS: ")
	fmt.Scan(&uas)
	

	var val = [3]int{
		kat,uts,uas,
	}
	fmt.Println(val)
	fmt.Println("---------------------------")
	fmt.Print("Rata-rata nilai = ")
	rataRata:= (6*val[0]/10)+(2*val[1]/10)+(2*val[2]/10)
	fmt.Println(rataRata)

	// manipulation
	fmt.Println(len(val))
	val[1]=80
	fmt.Println(val)
	fmt.Print("Rata-rata nilai = ")
	rataRata= (6*val[0]/10)+(2*val[1]/10)+(2*val[2]/10)
	fmt.Println(rataRata)
}