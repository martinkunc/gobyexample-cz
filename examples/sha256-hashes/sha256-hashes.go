// [_SHA256 hash funkce_](https://en.wikipedia.org/wiki/SHA-2) jsou
// často používané pro výpočet krátkých identit pro binární
// nebo textové bloby. Například TLS/SSL certifikáty používají
// SHA256 pro výpočet podpisu certifikátu. Tady je jak vypočítat
// SHA256 hashe v Go.

package main

// Go implementuje několik hash funkcí v různých balíčcích
// `crypto/*`.
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// Tady začínáme s novým hashem.
	h := sha256.New()

	// `Write` očekává bajty. Pokud máte řetězec `s`,
	// použijte `[]byte(s)` pro vynucení převodu do bajtů.
	h.Write([]byte(s))

	// Toto získá finalizovaný výsledek hashe jako výřez
	// bajtů. Parametr funkce `Sum` může být použitý pro přidání
	// k existujícímu výřezu bajtů: to obvykle není potřeba.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
