package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("--- Konverter Suhu ---")
	fmt.Print("Masukkan suhu dalam Celsius: ")

	// 1. Create a new reader to read from the standard input (keyboard).
	reader := bufio.NewReader(os.Stdin)

	// 2. Read the user's input until they press Enter.
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Gagal membaca input: %v", err)
	}

	// 3. Clean the input string and try to convert it to a float.
	input = strings.TrimSpace(input)
	suhu, err := strconv.ParseFloat(input, 64)
	if err != nil {
		// 4. If conversion fails, print an error and exit.
		//fmt.Print("Input Tidak Valid, hanya menerima angka\n")
		//log.Fatalf("Input Tidak Valid, hanya menerima angka")
		log.Fatalln("Input Tidak Valid, hanya menerima angka")
	}

	// 5. Perform the temperature conversions.
	reamur := suhu * 4 / 5
	fahrenheit := (suhu * 9 / 5) + 32
	//kelvin := suhu + 273.15

	// 6. Print the formatted results.
	fmt.Printf("Suhu dalam Reamur = %.2f\n", reamur)
	fmt.Printf("Suhu dalam Fahrenheit = %.2f\n", fahrenheit)
	//fmt.Printf("Suhu dalam Kelvin = %.2f\n", kelvin)
}
