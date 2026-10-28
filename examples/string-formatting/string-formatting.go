// Go nabízí výbornou podporu pro formátování řetězců
// v tradici `printf`. Zde jsou některé příklady
// běžných úkolů pro formátování řetězců.

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// Go nabízí několik "verbs" (sloves) navržených pro
	// formátování obecných Go hodnot. Například tohle vytiskne
	// instanci naší struktury `point`.
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// Pokud je hodnota struktura, varianta `%+v` bude obsahovat
	// také názvy prvků struktury.
	fmt.Printf("struct2: %+v\n", p)

	// Varianta `%#v` vypíše hodnotu reprezentovanou pomocí
	// Go syntaxe, tj. například kus zdrojového kódu
	// který by danou hodnotu vytvořil.
	fmt.Printf("struct3: %#v\n", p)

	// Pro vypsání typu hodnoty, použijte `%T`.
	fmt.Printf("type: %T\n", p)

	// Formátování boolean proměnných je přímočaré.
	fmt.Printf("bool: %t\n", true)

	// Je hodně možností jak formátovat integery.
	// Použij `%d` pro standardní formátování v desítkové soustavě.
	fmt.Printf("int: %d\n", 123)

	// Toto vypíše binární reprezentaci.
	fmt.Printf("bin: %b\n", 14)

	// Toto vypíše znak odpovídající danému integeru.
	fmt.Printf("char: %c\n", 33)

	// `%x` poskytuje hodnotu v hexadecimálním kódu.
	fmt.Printf("hex: %x\n", 456)

	// Máme také několik možností formátování pro
	// floaty. Pro základní desetinné formátování použijte `%f`.
	fmt.Printf("float1: %f\n", 78.9)

	// Slovesa `%e` a `%E` formátují floaty do (trochu odlišné
	// verzi) vědeckého zápisu.
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// Pro běžný tisk řetězců použijte `%s`.
	fmt.Printf("str1: %s\n", "\"string\"")

	// Pro tisk řetězce obaleného uvozovkami jako v Go kódu,
	// použijte `%q`.
	fmt.Printf("str2: %q\n", "\"string\"")

	// Jako jsme viděli už dříve s integery, `%x` renderuje
	// řetězec v 16-kové soustavě se dvěma výstupními znaky
	// pro jeden vstupní bajt.
	fmt.Printf("str3: %x\n", "hex this")

	// Pro výpis reprezentace ukazatele, použijte `%p`.
	fmt.Printf("pointer: %p\n", &p)

	// Při formátování čísel často budete chtít ovládat
	// šířku a přesnost výsledného čísla.
	// Pro specifikaci délku integeru, použije číslo
	// za slovesem `%`. Ve výchozím stavu bude výsledek
	// zarovnaný doprava a zleva doplněný mezerami.
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// Také můžete specifikovat šířku vypisovaných floatů,
	// ačkoliv obvykle také budete chtít omezit
	// desetinnou přesnost pomocí syntaxe se
	// zápisem šířka.přesnost
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// Pro levé zarovnání použij příznak `-`.
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Také možná budete chtít ovládat šířku při formátování
	// řetězců, obzvlášť pro dosažení jejích zarovnání v
	// tabulkovém výstupu. Pro základní pravé zarovnání na šířku.
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// Pro levé zarovnání použijte příznak `-` stejně jako u čísel.
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// Zatím jsme viděli `Printf`, který vypisuje
	// formátovaný řetězec na `os.Stdout`. `Sprintf` formátuje
	// a vrací řetězec, aniž by ho někam vypisoval.
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// Také můžete s pomocí `Fprintf` vypsat+formátovat i do
	// jiných `io.Writers` než jen `os.Stdout`.
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
