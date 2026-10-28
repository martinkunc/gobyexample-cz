// Go umožňuje _zotavit se_ (recover) z paniky pomocí
// zabudované funkce `recover`. `recover` dokáže zabránit
// tomu, aby `panic` ukončila program, a nechat ho místo
// toho pokračovat v provádění.

// Příklad, kde se to hodí: server by nechtěl spadnout,
// pokud v jednom z klientských spojení dojde ke kritické
// chybě. Místo toho by chtěl toto spojení zavřít a dál
// obsluhovat ostatní klienty. Přesně tohle ostatně ve
// výchozím stavu dělá `net/http` v Go pro HTTP servery.

package main

import "fmt"

// Tato funkce vyvolá paniku.
func mayPanic() {
	panic("problém")
}

func main() {
	// `recover` musí být zavolán uvnitř odložené funkce.
	// Když obklopující funkce vyvolá paniku, odložená
	// funkce se aktivuje a volání `recover` v ní paniku
	// zachytí.
	defer func() {
		if r := recover(); r != nil {
			// Návratová hodnota `recover` je chyba předaná
			// při volání `panic`.
			fmt.Println("Zotaveno. Chyba:\n", r)
		}
	}()

	mayPanic()

	// Tento kód se neprovede, protože `mayPanic` vyvolá
	// paniku. Provádění `main` se zastaví v místě paniky
	// a pokračuje v odložené uzávěře.
	fmt.Println("Po mayPanic()")
}
