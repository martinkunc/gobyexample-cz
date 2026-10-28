// _Mapy_ jsou [asociativní datové typy](http://en.wikipedia.org/wiki/Associative_array)
// zabudované do Go. (Někdy v jiných programovacích
// jazycích se jim také říká _hashe_ nebo _slovníky_).

package main

import "fmt"

func main() {

	// Pro vytvoření prázdné mapy, použijte zabudovaný
	// příkaz `make`:
	// `make(map[typ-klíče]typ-hodnoty)`.
	m := make(map[string]int)

	// Nastavte dvojice klíč/hodnota s použitím typického
	// zápisu
	// `jméno[klíč] = hodnota`
	m["k1"] = 7
	m["k2"] = 13

	// Vypsání např. s `fmt.Println` ukáže všechny
	// její dvojice klíč/hodnota.
	fmt.Println("mapa:", m)

	// Hodnotu můžeme získat pro daný klíč zápisem
	// `jméno[klíč]`.
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// Zabudovaná funkce `len` vrací při použití s mapou
	// počet dvojic klíč/hodnota
	fmt.Println("len:", len(m))

	// Zabudovaná funkce `delete` odstraňuje dvojice
	// klíč hodnota z mapy.
	delete(m, "k2")
	fmt.Println("mapa:", m)

	// Volitelná druhá návratová hodnota během získávání
	// hodnoty z mapy určuje jestli byl klíč v mapě
	// nalezen. Může být použita pro rozlišení mezi
	// chybějícími klíči a klíči s nulovými hodnotami
	// jako `0` nebo `""`. Zde nepotřebujeme vlastní
	// hodnotu, tak ji ignorujeme s použitím
	// _zahazovacího_ identifikátoru
	// `_`.
	_, nal := m["k2"]
	fmt.Println("nalezen:", nal)

	// Také můžete deklarovat a inicializovat novou mapu
	// na stejném řádku s takovouhle syntaxí:
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("mapa:", n)
}
