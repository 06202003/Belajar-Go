package main
import "fmt"

func main(){
	var val1,val2,presensi int16
	var lulusNilai, lulusPresensi, lulus bool

	fmt.Print("Total Presensi:")
	fmt.Scan(&presensi)
	fmt.Print("Nilai Pertama:")
	fmt.Scan(&val1)
	fmt.Print("Nilai Kedua:")
	fmt.Scan(&val2)

	lulusNilai = (val1+val2)/2 >=80
	lulusPresensi = presensi >=12
	fmt.Println(lulusNilai,lulusPresensi)
	lulus = lulusNilai && lulusPresensi
	fmt.Println(lulus)
}