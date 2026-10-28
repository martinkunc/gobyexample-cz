// Je možné definovat vlastní typy chyb, když k nim
// implementujeme metodu `Error()`. Zde je varianta
// příkladu nahoře, která používá vlastní typ pro
// konkrétní určení chyby argumentu.

package main

import (
	"errors"
	"fmt"
)

// Vlastní typ chyby má obvykle příponu "Error".
type argError struct {
	arg     int
	message string
}

// Přidáním této metody `Error` bude `argError`
// implementovat interface `error`.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {

		// Vrátíme naši vlastní chybu.
		return -1, &argError{arg, "s tímto nemohu pracovat"}
	}
	return arg + 3, nil
}

func main() {

	// `errors.AsType` je pokročilejší verze `errors.Is`.
	// Kontroluje, zda daná chyba (nebo jakákoli chyba
	// v jejím řetězci) odpovídá konkrétnímu typu chyby,
	// převede ji na hodnotu tohoto typu a vrátí také
	// `true`. Pokud shoda není, druhá návratová hodnota
	// je `false`.
	_, err := f(42)
	if ae, ok := errors.AsType[*argError](err); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err neodpovídá argError")
	}
}
