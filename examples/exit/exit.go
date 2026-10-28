// Použijte `os.Exit` pro okamžité ukončení s danou
// návratovou hodnotou.

package main

import "fmt"
import "os"

func main() {

	// Příkazy `defer` se při použití `os.Exit` _nespouští_
	// takže tento `fmt.Println` se nikdy nezavolá.
	defer fmt.Println("!")

	// Exit se stavem 3.
	os.Exit(3)
}

// Všimněte si, že naproti např. C, Go nepoužívá int
// návratovou hodnotu z `main` pro nastavení návratové
// hodnoty. Pokud budete potřebovat skončit s nenulovým
// návratovým kódem, použijte `os.Exit`.
