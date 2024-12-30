package main
import "fmt"

func main(){
	var nama string
	var nilai int
	
	fmt.Print("Nama: ")
	fmt.Scan(&nama)

	switch length:=len(nama); length>10{
		case true:
			fmt.Println("Tolong ubah jadi inisial")
		case false:
			switch nama {
			case "Yehezkiel":
				fmt.Println("Hallo",nama)
			case "David":
				fmt.Println("Hallo bro",nama)
			default:
				fmt.Println("Anda siapa?")
			}
	}

	fmt.Print("Nilai: ")
	fmt.Scan(&nilai)
	

	switch nilai >=0  && nilai<=100{
		case true:
			switch{
				case nilai >=80:
					fmt.Println("A")
				case nilai <80 && nilai>=70:
					fmt.Println("B")
				case nilai <70 && nilai>=60:
					fmt.Println("C")
				case nilai <60 && nilai>=50:
					fmt.Println("D")
				default:
				fmt.Println("E")
			}
		case false:
			fmt.Println("Nilai harus integer positif kurang dari 100")
	}


}