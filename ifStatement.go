package main
import "fmt"

func main(){
	var nama string
	var nilai int
	
	fmt.Print("Nama: ")
	fmt.Scan(&nama)

	if length:=len(nama); length>10{
		fmt.Println("Tolong ubah jadi inisial")
	}else{
		if nama == "Yehezkiel"{
			fmt.Println("Hallo",nama)
		}else if (nama =="David"){
			fmt.Println("Hallo bro",nama)
		}else{
			fmt.Println("Anda bukan Yehezkiel")
		}
	}

	fmt.Print("Nilai: ")
	fmt.Scan(&nilai)
	

	if nilai >=0  && nilai<=100{
		if nilai >=80 {
			fmt.Println("A")
		}else if nilai <80 && nilai>=70{
			fmt.Println("B")
		}else if nilai <70 && nilai>=60{
			fmt.Println("C")
		}else if nilai <60 && nilai>=50{
			fmt.Println("D")
		}else{
			fmt.Println("E")
		}
	}else{
		fmt.Println("Nilai harus integer positif kurang dari 100")
	}


}