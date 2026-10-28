// Řetězec je v Go výřez bajtů určený jen pro čtení. Jazyk
// i standardní knihovna zacházejí s řetězci zvláštním
// způsobem - jako s kontejnery textu zakódovaného
// v [UTF-8](https://cs.wikipedia.org/wiki/UTF-8).
// V jiných jazycích se řetězce skládají ze "znaků".
// V Go se koncept znaku nazývá `runa` (`rune`) - je to
// celé číslo, které reprezentuje Unicode kódový bod.
// [Tento blog post o Go](https://go.dev/blog/strings) je
// dobrým úvodem do tématu.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// `s` je `string`, kterému je přiřazena literálová
	// hodnota reprezentující slovo "ahoj" v thajštině.
	// Řetězcové literály jsou v Go text zakódovaný
	// v UTF-8.
	const s = "สวัสดี"

	// Protože řetězce jsou ekvivalentní s `[]byte`, toto
	// vrátí délku surových bajtů, které jsou v něm uložené.
	fmt.Println("Délka:", len(s))

	// Indexování řetězce vrací surové hodnoty bajtů na
	// každém indexu. Tento cyklus vypíše hexa hodnoty všech
	// bajtů, které tvoří kódové body v `s`.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Pro spočítání, kolik _run_ je v řetězci, můžeme použít
	// balíček `utf8`. Všimněte si, že doba běhu
	// `RuneCountInString` závisí na velikosti řetězce,
	// protože musí postupně dekódovat každou UTF-8 runu.
	// Některé thajské znaky jsou reprezentované UTF-8 kódovými
	// body, které mohou zabírat více bajtů, takže výsledek
	// tohoto počítání může být překvapivý.
	fmt.Println("Počet run:", utf8.RuneCountInString(s))

	// Cyklus `range` zachází s řetězci zvláštním způsobem
	// a dekóduje každou `runu` spolu s jejím posunem
	// v řetězci.
	for idx, runeValue := range s {
		fmt.Printf("%#U začíná na %d\n", runeValue, idx)
	}

	// Stejné iterace můžeme dosáhnout i explicitním
	// použitím funkce `utf8.DecodeRuneInString`.
	fmt.Println("\nPomocí DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U začíná na %d\n", runeValue, i)
		w = width

		// Tohle ukazuje předání hodnoty `rune` funkci.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	// Hodnoty uzavřené v jednoduchých uvozovkách jsou
	// _runové literály_. Hodnotu `rune` můžeme přímo
	// porovnat s runovým literálem.
	if r == 't' {
		fmt.Println("nalezeno tý")
	} else if r == 'ส' {
		fmt.Println("nalezeno so sua")
	}
}
