// [_SHA1 hash funkce_](http://en.wikipedia.org/wiki/SHA-1) jsou
// často používané pro výpočet krátkých identit pro binární
// nebo textové bloby. Například, [systém git pro správu verzí
// ](http://git-scm.com/) hojně používá SHA1 pro identifikování
// verzovaných souborů a adresářů. Tady je jak vypočítat
// SHA1 hashe v Go.

package main

// Go implementuje několik hash funkcí v různých balíčcích
// `crypto/*`.
import "crypto/sha1"
import "fmt"

func main() {
	s := "sha1 this string"

	// Vzor pro generování hashe je `sha1.New()`,
	// `sha1.Write(bytes)`, potom `sha1.Sum([]byte{})`.
	// Tady začínáme s novým hashem.
	h := sha1.New()

	// `Write` očekává bajty. Pokud máš řetězec `s`,
	// použij `[]byte(s)` pro vynucení převodu do bajtů.
	h.Write([]byte(s))

	// Toto získá finalizovaný výsledek hashe jako výřez
	// bajtů. Parametr funkce `Sum` může být použitý pro přidání
	// k existujícímu výřezů bajtů: to obvykle není potřeba.
	bs := h.Sum(nil)

	// Hodnoty SHA1 jsou často vypisované ve hexa, například
	// v git commitech. Použij formátovací verb `%x` pro převod
	// výsledku hashe na hexa string.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
