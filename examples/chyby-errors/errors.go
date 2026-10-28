// V Go je přirozené komunikovat chyby pomocí
// explicitních, oddělených návratových hodnot. To je
// v kontrastu k výjimkám používaných v jazycích jako Java
// a Ruby a přetíženým samostatným hodnotám výsledku/chyby
// někdy používané v C. Go přístup usnadňuje zjištění zda
// funkce vrací chyby a jejich ošetření stejnými
// jazykovými prostředky jako pro úkoly které se
// netýkají chyb.

package main

import "errors"
import "fmt"

// Podle konvence, chyby jsou poslední vrácenou hodnotou
// a mají typ `error`, což je zabudovaný interface.
func f1(arg int) (int, error) {
	if arg == 42 {

		// `errors.New` vytváří základní hodnotu `error`
		// se zadanou chybovou hláškou.
		return -1, errors.New("nemohu pracovat se 42")

	}

	// Hodnota `nil` na pozici chyby naznačuje že nedošlo
	// k chybě.
	return arg + 3, nil
}

// Je možné použít vlastní typy jako `error`, když k nim
// implementujeme metodu `Error()`. Zde je varianta
// příkladu nahoře která používá vlastní typ pro
// konkrétní určení chyby argumentu.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// V tomto případě používáme zápis `&argError` pro
		// novou strukturu, a dodáváme ji hodnoty pro
		// dva prvky struktury `arg` a `prob`.
		return -1, &argError{arg,
			"s tímto nemohu pracovat"}
	}
	return arg + 3, nil
}

func main() {

	// Tyto dva cykly testují každou z našich funkcí
	// vracejících chyby.
	// Zapamatujte si, že použití kontroly chyb uvnitř
	// řádku je běžný idiom v Go.
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 selhala:", e)
		} else {
			fmt.Println("f1 fungovala:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 selhala:", e)
		} else {
			fmt.Println("f2 fungovala:", r)
		}
	}

	// Pokud budete potřebovat v programu přistoupit k
	// datům ve vlastní chybě, budete muset získat
	// instanci vlastního typu chyby pomocí typového
	// vynucení (type assertion).
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
