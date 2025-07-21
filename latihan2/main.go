package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Pengukuran menampung data lokasi, suhu, dan klasifikasinya.
type Pengukuran struct {
	Lokasi      string
	Celsius     float64
	Klasifikasi string
}

// KlasifikasikanSuhu adalah method untuk menentukan klasifikasi suhu
// dan menyimpannya ke dalam field Klasifikasi.
func (p *Pengukuran) KlasifikasikanSuhu() {
	if p.Celsius < 18 {
		p.Klasifikasi = "Dingin"
	} else if p.Celsius >= 18 && p.Celsius <= 25 {
		p.Klasifikasi = "Hangat"
	} else {
		p.Klasifikasi = "Panas"
	}
}

func main() {

	// Atur logger untuk tidak menampilkan tanggal dan waktu
	log.SetFlags(0)
	fmt.Println("--- Konverter & Klasifikasi Suhu ---")
	reader := bufio.NewReader(os.Stdin)

	// 1. Meminta dan memvalidasi input lokasi
	fmt.Print("Masukkan lokasi pengukuran suhu: ")
	lokasiInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Gagal membaca input lokasi: %v", err)
	}
	lokasiInput = strings.TrimSpace(lokasiInput)

	// Validasi: Lokasi tidak boleh kosong dan tidak boleh angka
	if lokasiInput == "" {
		log.Fatalln("Lokasi tidak boleh kosong.")
	}
	if _, err := strconv.ParseFloat(lokasiInput, 64); err == nil {
		log.Fatalln("Input Lokasi tidak boleh hanya berupa angka.")
	}

	// 2. Meminta dan memvalidasi input suhu
	fmt.Print("Masukkan suhu dalam Celsius: ")
	suhuInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Gagal membaca input suhu: %v", err)
	}
	suhuInput = strings.TrimSpace(suhuInput)
	suhuCelsius, err := strconv.ParseFloat(suhuInput, 64)
	if err != nil {
		log.Fatalln("Input Suhu Tidak Valid, hanya menerima angka.")
	}

	// 3. Buat instance struct dan isi datanya
	dataPengukuran := Pengukuran{
		Lokasi:  lokasiInput,
		Celsius: suhuCelsius,
	}

	// 4. Panggil method untuk klasifikasi suhu
	dataPengukuran.KlasifikasikanSuhu()

	// 5. Lakukan konversi suhu lainnya
	reamur := dataPengukuran.Celsius * 4 / 5
	fahrenheit := (dataPengukuran.Celsius * 9 / 5) + 32

	// 6. Tampilkan semua hasil
	//fmt.Println("\n--- Hasil Pengukuran ---")
	fmt.Printf("Suhu di %s adalah %s\n", dataPengukuran.Lokasi, dataPengukuran.Klasifikasi)
	fmt.Printf("Suhu di %s  dalam Reamur =  %.0f\n", dataPengukuran.Lokasi, reamur)
	fmt.Printf("Suhu di %s  dalam Fahrenheit =  %.0f\n", dataPengukuran.Lokasi, fahrenheit)

	//fmt.Printf("Klasifikasi Suhu: %s\n", dataPengukuran.Klasifikasi)
	//fmt.Printf("Suhu dalam Reamur = %.2f\n", reamur)
	//fmt.Printf("Suhu dalam Fahrenheit = %.2f\n", fahrenheit)
}
