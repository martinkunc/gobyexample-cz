// Pole je v Go očíslovanou sekvencí prvků předem
// dané délky

package main

import "fmt"

func main() {

	// Zde vytváříme pole `a` které bude mít přesně 5
	// `int`ů. Typ prvků a délka jsou obě součást typu
	// pole. Přirozeně, pole je pole nulových hodnot,
	// což pro pole `int`ů znamená pole naplněné na `0`.
	var a [5]int
	fmt.Println("prázdné:", a)

	// Můžeme nastavit hodnotu na indexu použitím syntaxe
	// `array[index] = value` a získat hodnotu s použitím
	// `array[index]`.
	a[4] = 100
	fmt.Println("nastaveno:", a)
	fmt.Println("získáno:", a[4])

	// Zabudovaná funkce `len` vrací délku pole
	fmt.Println("délka:", len(a))

	// Použij tento zápis pro deklaraci a inicializaci
	// pole v jednom řádku.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("deklarované:", b)

	// Typy polí jsou jedno-rozměrné, ale můžete typy
	// skládat a vytvořit tak vícerozměrné datové
	// struktury.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
