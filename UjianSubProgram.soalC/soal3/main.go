//Randy Lambas Batubara
//2311102096
package main

import (
	"fmt"
)


func divide(n, m int) (int, int) {
	if n < m {
		return 0, n 
	}
	quotient, remainder := divide(n-m, m)
	return quotient + 1, remainder 
}

func main() {
	var n, m int

	fmt.Print("Randy Lambas Batubara\n")
	fmt.Print("2311102096\n")
	fmt.Print("Masukkan bilangan bulat positif n (pembilang): ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan bulat positif m (penyebut): ")
	fmt.Scan(&m)

	
	if n <= 0 || m <= 0 {
		fmt.Println("Harap masukkan bilangan bulat positif.")
		return
	}
	quotient, remainder := divide(n, m)
    fmt.Printf("Hasil pembagian %d dengan %d adalah: %d\n", n, m, quotient)
	fmt.Printf("Sisa dari pembagian %d dengan %d adalah: %d\n", n, m, remainder)
}