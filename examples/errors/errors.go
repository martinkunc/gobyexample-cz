// V Go je přirozené komunikovat chyby pomocí
// explicitních, oddělených návratových hodnot. To je
// v kontrastu k výjimkám používaným v jazycích jako Java,
// Python a Ruby a přetíženým samostatným hodnotám
// výsledku/chyby někdy používaným v C. Go přístup
// usnadňuje zjištění, zda funkce vrací chyby, a jejich
// ošetření stejnými jazykovými prostředky jako pro úkoly,
// které se netýkají chyb.
//
// Pro další podrobnosti se podívejte na dokumentaci
// [balíčku errors](https://pkg.go.dev/errors)
// a [tento blog post](https://go.dev/blog/go1.13-errors).

package main

import (
	"errors"
	"fmt"
)

// Podle konvence, chyby jsou poslední vrácenou hodnotou
// a mají typ `error`, což je zabudovaný interface.
func f(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` vytváří základní hodnotu `error`
		// se zadanou chybovou hláškou.
		return -1, errors.New("nemohu pracovat se 42")
	}

	// Hodnota `nil` na pozici chyby naznačuje že nedošlo
	// k chybě.
	return arg + 3, nil
}

// Sentinelová chyba (sentinel error) je předem
// deklarovaná proměnná, která označuje konkrétní
// chybový stav.
var ErrDoselCaj = errors.New("čaj došel")
var ErrProud = errors.New("nelze uvařit vodu")

func uvarCaj(arg int) error {
	if arg == 2 {
		return ErrDoselCaj
	} else if arg == 4 {

		// Chyby můžeme obalovat chybami vyšší úrovně
		// a přidat tak kontext. Nejjednodušší způsob je
		// sloveso `%w` ve `fmt.Errorf`. Obalené chyby
		// vytvářejí logický řetězec (A obaluje B, které
		// obaluje C, atd.), na který se lze dotazovat
		// funkcemi jako `errors.Is` a `errors.AsType`.
		return fmt.Errorf("vaření čaje: %w", ErrProud)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// Použití kontroly chyb přímo na řádku s `if`
		// je běžný idiom v Go.
		if r, e := f(i); e != nil {
			fmt.Println("f selhala:", e)
		} else {
			fmt.Println("f fungovala:", r)
		}
	}

	for i := range 5 {
		if err := uvarCaj(i); err != nil {

			// `errors.Is` kontroluje, zda daná chyba (nebo
			// jakákoli chyba v jejím řetězci) odpovídá
			// konkrétní hodnotě chyby. To je užitečné
			// hlavně u obalených nebo vnořených chyb, kde
			// tak můžete v řetězci chyb rozpoznat konkrétní
			// typy chyb nebo sentinelové chyby.
			if errors.Is(err, ErrDoselCaj) {
				fmt.Println("Měli bychom koupit nový čaj!")
			} else if errors.Is(err, ErrProud) {
				fmt.Println("Teď je tma.")
			} else {
				fmt.Printf("neznámá chyba: %s\n", err)
			}
			continue
		}

		fmt.Println("Čaj je hotový!")
	}
}
