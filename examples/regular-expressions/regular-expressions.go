// Go nabízí zabudovanou podporu pro [regulární výrazy](http://en.wikipedia.org/wiki/Regular_expression).
// Tady je několik příkladů běžných úkolů které se vztahují k regulárním výrazům v Go

package main

import "bytes"
import "fmt"
import "regexp"

func main() {

	// Toto otestuje jestli vzorek regulárního výrazu odpovídá danému řetězci.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Výše jsme použili řetězec regulárního vzoru přímo, ale pro další
	// úlohy s metodami z balíku regexp je bude potřeba `Kompilovat` do
	// optimalizované struktury `Regexp`.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// V této struktuře je dostupné množství metod. Tady je
	// test zda regexp odpovídá (match) textu jako jsme viděli výše.
	fmt.Println(r.MatchString("peach"))

	// Toto najde odpovídající shodu pro regexp.
	fmt.Println(r.FindString("peach punch"))

	// Tohle také najde první odpovídající shodu, ale vrátí
	// počáteční a koncové indexy pro shodu místo shodného textu.
	fmt.Println(r.FindStringIndex("peach punch"))

	// Varianty `Submatch` poskytují dvě informace,
	// jak shodu celého vzorku tak i skupiny
	// shodující se uvnitř. Například tohle vrátí informace
	// o obou `p([a-z]+)ch` i `([a-z]+)`.
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Podobně tohle vrátí informace o indexech shody
	// i skupinách.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// Varianty `All` těchto funkcí platí pro všechny
	// schody ve vstupu, ne jenom k první. Například
	// pro nalezení všech shod pro regexp.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// Varianty s `All` jsou také dostupné i pro další
	// funkce které jsme viděli výše.
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Předáním kladného čísla v druhém argumentu u těchto
	// funkcí omezí počet vrácených shod.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Naše příklady výše měly řetězcové argumenty a jmenovaly
	// se například `MatchString`. Také jim můžeme
	// poskytnout argumenty pomocí `[]byte` a vypustit z názvu
	// funkce `String`.
	fmt.Println(r.Match([]byte("peach")))

	// Když vytváříme konstanty s regulárními výrazy
	// můžeme použít kromě `Compile` i `MustCompile`.
	// Čisté `Compile` nebude fungovat u konstant
	// protože má 2 návratové hodnoty.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// Balíček `regexp` také může být použitý k nahrazování
	// podmnožin řetězců jinými hodnotami.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Varianta s `Func` umožňuje transformovat nalezený
	// text pomocí dané funkce.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
