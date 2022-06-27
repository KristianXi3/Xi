package main
import "fmt"

func main () {
	nama := [] string{"Budi1","Budi2","Budi3", "Budi4"}
	for i :=0 ; i < len(nama); i++ {
		fmt.Println (nama[i])
	}
}