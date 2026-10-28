// _Pole_ je v Go očíslovanou sekvencí prvků předem
// dané délky. V typickém Go kódu jsou daleko běžnější
// [výřezy](slices); pole jsou užitečná v některých
// speciálních situacích.

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

	// Počet prvků za vás může spočítat i překladač,
	// pomocí `...`.
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("deklarované:", b)

	// Pokud zadáte index pomocí `:`, prvky mezi nimi
	// budou vynulované.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("index:", b)

	// Typy polí jsou jedno-rozměrné, ale můžete typy
	// skládat a vytvořit tak vícerozměrné datové
	// struktury.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Vícerozměrná pole je také možné vytvořit
	// a inicializovat najednou.
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
