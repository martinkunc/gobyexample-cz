// _Výřezy_, v angličtině Slices, jsou důležitým datovým
// typem v Go, poskytujícím sekvencím bohatší rozhraní
// než polím.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Na rozdíl od polí, typ výřezů je jen typ jejich
	// prvků (ne počet prvků).
	// Neinicializovaný výřez je roven nil a má
	// délku 0.
	var s []string
	fmt.Println("neinic:", s, s == nil, len(s) == 0)

	// Pro vytvoření výřezu nenulové délky
	// použijte zabudovanou funkci `make`. Tady vytváříme
	// výřez `string`ů délky `3` (automaticky
	// inicializovaný na nulové hodnoty).
	// Ve výchozím stavu je kapacita nového výřezu rovna
	// jeho délce; pokud předem víme, že výřez bude růst,
	// můžeme kapacitu předat explicitně jako další
	// parametr funkce `make`.
	s = make([]string, 3)
	fmt.Println("prázdné:", s, "délka:", len(s), "kap.:", cap(s))

	// Můžeme nastavovat a získávat hodnoty jako u polí
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("nastavené:", s)
	fmt.Println("získáno:", s[2])

	// `len` vrací podle očekávání délku výřezu.
	fmt.Println("délka:", len(s))

	// Navíc k temto základním operacím podporují výřezy
	// několik dalších, díky nimž jsou bohatší než pole.
	// Jednou je zabudovaný `append`, který vrací výřez
	// obsahující jednu nebo více nových hodnot.
	// Všimněte si, že musíme přijmout návratovou hodnotu
	// `append`u, protože můžeme dostat novou hodnotu
	// výřezu.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("přidané:", s)

	// Výřezy také mohou být kopírované pomocí `copy`'.
	// Zde vytváříme prázdný výřez `c` o stejné délce jako
	// `s` a kopírujeme do `c` z `s`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("kopie:", c)

	// Výřezy podporují "výřezový" operátor pomocí zápisu
	// `vyrez[dolni:horni]`. Například, toto vytvoří výřez
	// prvků `s[2]`, `s[3]`, a `s[4]`.
	l := s[2:5]
	fmt.Println("vy1:", l)

	// Toto vyřízne od začátku až do (ale nezahrnuje)
	// `s[5]`.
	l = s[:5]
	fmt.Println("vy2:", l)

	// A tohle vyřízne od (a včetně) `s[2]` do konce.
	l = s[2:]
	fmt.Println("vy3:", l)

	// Také můžeme deklarovat a inicializovat proměnnou
	// s výřezem na jednom řádku.
	t := []string{"g", "h", "i"}
	fmt.Println("dekl:", t)

	// Balíček `slices` obsahuje řadu užitečných pomocných
	// funkcí pro práci s výřezy.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Výřezy se mohou skládat do vícerozměrných datových
	// struktur. Délky vnitřních výřezů se mohou měnit,
	// na rozdíl od více-rozměrných polí.
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
