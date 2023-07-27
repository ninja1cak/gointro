package main

import "fmt"

const firstName string = "Hauzan "
const lastName string = "khoirullah"

func age(yearNow int, yearBirth int) int {
	substract := 0
	substract = yearNow - yearBirth
	return substract
}

// func main() {
// 	var angka [4]int

// 	mapExamps()
// 	angka = [4]int{
// 		1,
// 		2,
// 		3,
// 		4,
// 	}
// 	fmt.Println(angka[2])

// 	var fruits = [4]string{
// 		"apel",
// 		"jeruk",
// 		"duren",
// 		"manggis",
// 	}

// 	var fruitsSlice = []string{
// 		"apel",
// 		"jeruk",
// 		"duren",
// 		"manggis",
// 		"anggur",
// 	}

// 	fruits_1 := fruitsSlice[0:3]

// 	fmt.Println(len(fruitsSlice), "slice")
// 	fmt.Println(fruits_1, "slice")

// 	for index, value := range fruits {
// 		fmt.Println(index, value)
// 	}

// 	// fmt.Println("hello word, total = ", age(2023, 2000))
// 	// for i := 0; i < 5; i++ {
// 	// 	if i%2 == 0 {
// 	// 		fmt.Println(i, "fazz")

// 	// 	}
// 	// }

// 	// var i int = 0
// 	// for i < 10 {
// 	// 	fmt.Println(i)
// 	// 	i++

// 	// }
// }

func mapExamps() {
	//{key: "value"}
	var angka = map[string]int{
		"satu": 1,
		"dua":  2,
	}

	var talent = []map[string]string{
		{"satu": "1",
			"dua": "2"},
		{
			"tiga":  "3",
			"empat": "4",
		},
	}
	fmt.Println("object", angka["satu"])
	fmt.Println("object", talent)

	for key, value := range angka {
		fmt.Println(key, "=", value)
	}
}
