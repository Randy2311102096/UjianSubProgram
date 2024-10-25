//Randy Lambas Batubara
//2311102096
package main

import (
	"fmt"
	"strconv"
)

func potongBilangan(bilangan int) {
	strBilangan_2311102096 := strconv.Itoa(bilangan)
	panjang := len(strBilangan_2311102096)

	var mid int
	if panjang%2 == 0 {
		mid = panjang / 2
	} else {
		mid = panjang/2 + 1
	}

	bagian1, err1 := strconv.Atoi(strBilangan_2311102096[:mid])
	bagian2, err2 := strconv.Atoi(strBilangan_2311102096[mid:])

	if err1 != nil || err2 != nil {
		fmt.Println("Terjadi kesalahan saat mengonversi bagian bilangan.")
		return
	}

	hasilPembagian := bagian1 / bagian2

	fmt.Printf("Bilangan 1: %d\n", bagian1)
	fmt.Printf("Bilangan 2: %d\n", bagian2)
	fmt.Printf("Hasil pembagian: %d\n", hasilPembagian)
}

func main() {
	nim := 2311102096

	var bilangan int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	_, err := fmt.Scan(&bilangan)

	if err != nil || bilangan <= 10 {
		fmt.Println("Input tidak valid. Masukkan bilangan bulat positif lebih besar dari 10.")
		return
	}

	potongBilangan(bilangan)

	fmt.Println("NIM:", nim)
}
