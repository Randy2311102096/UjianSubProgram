//Randy Lambas Batubara
//2311102096
package main

import (
	"fmt"
	"strconv"
)

// digit sama
func SDGS(n int) bool {
	str := strconv.Itoa(n)
	for i := 1; i < len(str); i++ {
		if str[i] != str[0] {
			return false
		}
	}
	return true
}

// beda digit
func SDB(n int) bool {
	str := strconv.Itoa(n)
	digitMap := make(map[rune]bool)
	for _, ch := range str {
		if digitMap[ch] {
			return false
		}
		digitMap[ch] = true
	}
	return true
}

func main() {
	var N int
	fmt.Print("Randy Lambas Batubara\n")
	fmt.Print("2311102096\n\n")
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&N)

	hadiahUtama := 0
	hadiahSembako := 0
	hadiahKonsol := 0

	for i := 1; i <= N; i++ {
		var nomorKartu int
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&nomorKartu)

		if SDGS(nomorKartu) {
			fmt.Println("Hadiah Utama")
			hadiahUtama++
		} else if SDB(nomorKartu) {
			fmt.Println("Hadiah Sembako")
			hadiahSembako++
		} else {
			fmt.Println("Hadiah Konsol")
			hadiahKonsol++
		}
	}

	fmt.Printf("Jumlah yang memperoleh Hadiah Utama: %d\n", hadiahUtama)
	fmt.Printf("Jumlah yang memperoleh Hadiah Sembako: %d\n", hadiahSembako)
	fmt.Printf("Jumlah yang memperoleh Hadiah Konsol: %d\n", hadiahKonsol)
}
