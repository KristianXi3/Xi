package main

import "fmt"

func main(){
	x := 0
	for  x <= 10{
		if (x % 2 == 0){
			fmt.Println("Genap")
		} else {
			fmt.Println("Ganjil")
		}
		x++
	}

}