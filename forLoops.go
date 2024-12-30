package main
import "fmt"

func main(){
	counter:=1

	for counter<=10{
		fmt.Println(counter)
		counter++
	}

	fmt.Println("==================================================")
	for i:=0;i<10;i++{
		fmt.Println(i)
	}

	datas:= []string{
		"Yehezkiel", "David", "Setiawan",
	}

	fmt.Println("==================================================")
	for i:=0;i<len(datas);i++{
		fmt.Println(datas[i])
	}

	datas = append(datas,"Ganteng")

	fmt.Println("==================================================")
	for i:=0;i<len(datas);i++{
		fmt.Println(datas[i])
	}

	fmt.Println("======================Index============================")
	for index,data := range datas{
		fmt.Println("Nama ke-",index+1,"adalah",data)
	}

	fmt.Println("===================None Index========================")
	for _,data := range datas{
		fmt.Println("Nama: ",data)
	}

	fmt.Println("===================Break========================")
	for _,data := range datas{
		if data == "David"{
			break
		}else{
			fmt.Println("Nama: ",data)
		}

	}

	fmt.Println("========================Continue==========================")
	for j:=0;j<10;j++{
		if j %2==0{
			continue
		}
		fmt.Println(j)
		

	}

}