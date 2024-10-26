package main

import "fmt"

func main() {
	var jumlahMahasiswa int

	fmt.Printf("Masukkan jumlah mahasiswa: ")
	fmt.Scanln(&jumlahMahasiswa)

	var dataMahasiswa = []struct {
		nama string
		nilai [5]int
	}{}

	for i := 1; i <= jumlahMahasiswa; i++ {
		var nama string
		var nilai [5]int

		fmt.Printf("Masukkan nama mahasiswa ke-%d: ", i)
		fmt.Scanln(&nama)

		fmt.Printf("Masukkan nilai untuk 5 mata kuliah (pisahkan dengan spasi): ")
		fmt.Scanf("%d %d %d %d %d\n", &nilai[0], &nilai[1], &nilai[2], &nilai[3], &nilai[4])

		dataMahasiswa = append(dataMahasiswa, struct {
			nama string
			nilai [5]int
		}{nama, nilai})
	}

	var mahasiswaTerbaik struct {
		nama string
		rataRata float32
	}

	for _, mahasiswa := range dataMahasiswa {
		var totalNilai int
		for _, nilai := range mahasiswa.nilai {
			totalNilai += nilai
		}

		var rataRata = float32(totalNilai) / 5;
		grade := ""
		
		if rataRata >= 80 {
			grade = "A"
		}
		if rataRata >= 70 && rataRata < 80 {
			grade = "B"
		}
		if rataRata >= 60 && rataRata < 70 {
			grade = "C"
		}
		if rataRata >= 50 && rataRata < 60 {
			grade = "D"
		}
		if rataRata < 50 {
			grade = "E"
		}

		if rataRata > mahasiswaTerbaik.rataRata {
			mahasiswaTerbaik.nama = mahasiswa.nama
			mahasiswaTerbaik.rataRata = rataRata
		}

		fmt.Printf("Nama: %s, Rata-rata: %.2f, Nilai Huruf: %s", mahasiswa.nama, rataRata, grade)
		fmt.Printf("\n\n")
	}

	fmt.Printf("Mahasiswa dengan nilai rata-rata tertinggi: %s dengan rata-rata %.2f", mahasiswaTerbaik.nama, mahasiswaTerbaik.rataRata)
}