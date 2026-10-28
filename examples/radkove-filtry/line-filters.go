// _Řádkový filtr_ je běžný druh programu který čte
// vstup ze stdin, zpracovává ho, a pak vypisuje nějaký
// odvezený výstup na stdout. Běžné řádkové filtry
// jsou `grep` a `sed`.

// Zde je příklad řádkového filtru v Go který zapisuje
// kapitalizovanou verzi každého textu na vstupu.
// Můžeš použít tento vzorek při psaní svých vlastních
// řádkových filtrů.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Zabalení nebufferovaného `os.Stdin` do bufferovaného
	// skenneru nám poskytne užitečnou metodu `Scan`
	// která posouvá scanner na další token; kterým je
	// další řádek ve výchozím skeneru.
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text` vrací aktuální token, tady další řádek
		// ze vstupu.
		ucl := strings.ToUpper(scanner.Text())

		// Zapiš kapitalizovaný řádek.
		fmt.Println(ucl)
	}

	// Zkontroluj chyby během `Scan`. Konec souboru je
	// očekávaný a není hlášený ve `Scan` jako chyba.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "chyba:", err)
		os.Exit(1)
	}
}
