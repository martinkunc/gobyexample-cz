// [_Příznaky na příkazové řádce_](http://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// jsou běžným způsobem pro nastavení voleb pro příkazovou řádku
// programů. Například, u `wc -l` volba `-l` je příznak
// příkazové řádky (command-line flag).

package main

// Go poskytuje balíček `flag` pro podporu základního
// parsování příznaků z příkazové řádky. Použijeme
// tento balíček pro implementaci našeho příkladu
// programu pro příkazovou řádku.
import "flag"
import "fmt"

func main() {

	// Základní deklarace příznaků jsou dostupné pro
	// volby typu řetězec, integer, a boolean. Zde
	// deklarujeme řetězcový příznak `word` s výchozí
	// hodnotou `"foo"` a krátkým popisem. Tato funkce
	// `flag.String` vrací ukazatel na řetězec ;
	// (ne hodnotu řetězce) a dále uvidíme jak těnto
	// ukazatel použít.
	wordPtr := flag.String("word", "foo", "a string")

	// Toto deklaruje příznaky `numb` a `fork` s použitím
	// podobného přístupu jako u příznaku `word`.
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// Také je možné deklarovat příznak který používá
	// existující proměnnou deklarovanou jinde v
	// programu. Všimněte si, že do funkce deklarující
	// příznak musíme předat ukazatel.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Když jsou všechny příznaky deklarované, zavoláme `flag.Parse()`
	// pro naparsování příkazové řádky.
	flag.Parse()

	// Tady jenom vypíšeme naparsované volby a
	// všechny zbývající poziční argumenty. Všimněte si že
	// ukazatele musíme dereferencovat pomocí např.
	// `*wordPtr` k získání vlastních hodnot voleb.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
