package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	var (
		length               = 0
		totalPerkalian       = 1
		total, nilaiTengah   int
		nilaiRataRata, value float64
	)

	// input the length of input
	fmt.Println("Enter the length of inputs: ")
	_, err := fmt.Scanln(&length)
	if err != nil {
		log.Fatal(errors.New("Please input integer value"))
		return
	}

	// input the number
	fmt.Println("Enter an integer value : ")
	numbers := make([]int, length)

	// insert the number into array
	for i := 0; i < length; i++ {
		_, err := fmt.Scanln(&numbers[i])
		if err != nil {
			log.Fatal(errors.New("Please input integer value"))
			return
		}
	}

	// sort number by ascending (small to large)
	for i, num := range numbers {
		for j, n := range numbers {
			if n > num {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}

	// count nilai rata rata
	for _, nums := range numbers {
		total += nums
	}

	nilaiRataRata = float64(total / len(numbers))

	// mencari nilai tengah
	value = float64(len(numbers) / 2)
	nilaiTengah = numbers[int(math.Round(value))]

	// mencari total perkalian
	for _, nums := range numbers {
		totalPerkalian = totalPerkalian * nums
	}

	fmt.Println("RESULT")
	fmt.Println("Urutan dari yang terkecil : ", numbers)
	fmt.Println("Nilai Rata-rata : ", nilaiRataRata)
	fmt.Println("Nilai Tengah : ", nilaiTengah)
	fmt.Println("Hasil Perkalian : ", totalPerkalian)
}
