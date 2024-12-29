package main
import "fmt"

func main(){
	var nama,nrp string

	fmt.Print("Nama :")
	fmt.Scan(&nama)
	fmt.Print("NRP :")
	fmt.Scan(&nrp)
	fmt.Println(nama,"-",nrp)
}