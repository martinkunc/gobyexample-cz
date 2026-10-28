// Go nabízí výbornou podporu pro formátování řetězců
// v tradici `printf`. Zde jsou některé příklady
// běžných úkolů pro formátování řetězců.

package main

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {

	// Go nabízí několik "verbs" (sloves) navržených pro
	// formátování obecných Go hodnot. Například tohle vytiskne
	// instanci naši struktury `point`.
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// Pokud je hodnota struktura, varianta `%+v` bude obsahovat
	// také názvy prvků struktury.
	fmt.Printf("%+v\n", p)

	// Varianta `%#v` vypíše hodnotu reprezentovanou pomocí
	// Go syntaxe, tj. například kus zdrojového kódu
	// který by danou hodnotu vytvořil.
	fmt.Printf("%#v\n", p)

	// Pro vypsání typu hodnoty, použijte `%T`.
	fmt.Printf("%T\n", p)

	// Formátování boolean proměnných je přímočaré.
	fmt.Printf("%t\n", true)

	// Je hodně možností jak formátovat integery.
	// Použij `%d` pro standardní formátování v desítkové soustavě.
	fmt.Printf("%d\n", 123)

	// Toto vypíše binární reprezentaci.
	fmt.Printf("%b\n", 14)

	// Toto vypíše znak odpovídající danému integeru.
	fmt.Printf("%c\n", 33)

	// `%x` poskytuje hodnotu v hexadecimálním kódu.
	fmt.Printf("%x\n", 456)

	// Máme také několik možností formátování pro
	// floaty. Pro základní desetinné formátování použijte `%f`.
	fmt.Printf("%f\n", 78.9)

	// Slovesa `%e` a `%E` formátují floaty do (trochu odlišné
	// verzi) vědeckého zápisu.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// Pro běžný tisk řetězců použijte `%s`.
	fmt.Printf("%s\n", "\"string\"")

	// Pro tisk řetězce obaleného uvozovkami jako v Go kódu,
	// použijte `%q`.
	fmt.Printf("%q\n", "\"string\"")

	// Jako jsme viděli už dříve s integery, `%x` renderuje
	// řetězec v 16-kové soustavě se dvěma výstupními znaky
	// pro jeden vstupní bajt.
	fmt.Printf("%x\n", "hex this")

	// Pro výpis reprezentace ukazatele, použijte `%p`.
	fmt.Printf("%p\n", &p)

	// Při formátování čísel často budete chtít ovládat
	// šířku a přesnost výsledného čísla.
	// Pro specifikaci délku integeru, použije číslo
	// za slovesem `%`. Ve výchozím stavu bude výsledek
	// zarovnaný doprava a zleva doplněný mezerami.
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// Také můžete specifikovat šířku vypisovaných floatů,
	// ačkoliv obvzkle také budete chtít omezit
	// desetinnou přesnost pomocí syntaxe se
	// zápisem šířka.přesnost
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// Pro levé zarovnání použij příznak `-`.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Také možná budete chtít ovládat šířku při formátování
	// řetězců, obzvlášť pro dosažení jejích zarovnání v
	// tavulkovém výstupu. Pro základní pravé zarovnání na šířku.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// Pro levé zarovnání použitje příznak  `-` stejně jako u čísel.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// Zatím jsme viděli `Printf`, který vupisuje
	// formátovaný řetězec na `os.Stdout`. `Sprintf` formátuje
	// a vrací řezězec bež toho aby ho někam vypisoval..
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Také můžete s pomocí `Fprintf` vypsat+formátovat i do
	// jiných `io.Writers` než jen `os.Stdout`.
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
