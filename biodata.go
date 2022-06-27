package main 

import ("fmt" 
"strconv"
 "os"
)

type Absen struct {
	Nama string
	Alamat string
	Pekerjaan string
	Alasan string
}

func main () {
	var list = []Absen{
		{ Nama : "Budi1",Alamat : "Jakarta",Pekerjaan : "Karyawan", Alasan : "Kerja",},
		{ Nama : "Budi2",Alamat : "Jakarta",Pekerjaan : "Pegawai",	Alasan : "Kerja",},
		{ Nama : "Budi3",Alamat : "Jakarta",Pekerjaan : "Pekerja",	Alasan : "Kerja",},
	}
	

	// for _, i := range list{
	// 	fmt.Println("%s\n%s\n%s\n%s\n", v.Nama, v.Alamat, v.Pekerjaan, v.Alasan)
	// }
	i, _ := strconv.Atoi(os.Args[1])
	fmt.Println(list[i-1])

}

