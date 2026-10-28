// Parsování čísel z řetězců je základní ale běžnou
// základní úlohou v mnoha programech; tady je jak
// to dělat v Go.

package main

// Zabudovaný balíček `strconv` poskytuje parsování čísel.
import "strconv"
import "fmt"

func main() {

	// S `ParseFloat`, těch `64` říká kolik bitů přesnosti
	// parsovat.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// U `ParseInt`, ta `0` znamená že se má odvodit základ
	// z řetězce. `64` vyžaduje aby se výsledek vešel do 64
	// bitů.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` rozpozná hexa formátovaná čísla.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// Také je dostupná `ParseUint`.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` je funkce pro pohodlí pro základní parsování `int`u
	// v desítkové soustavě.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Funkce Parse vrátí chybu u špatného vstupu.
	_, e := strconv.Atoi("co?")
	fmt.Println(e)
}
