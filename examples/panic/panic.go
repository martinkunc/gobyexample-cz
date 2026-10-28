// Význam klíčového slova `panic` je že se stalo něco neočekávaně
// špatného. Většinou ho používáme pro rychlé selhání během chyb
// které by se během normálního běhu neměly stát a nebo které
// nejsme připraveni elegantně ošetřit.

package main

import "os"

func main() {

	// Na této stránce použijeme panic pro kontrolu nečekaných
	// chyb. Tohle je jediný program na stránce, který je
	// navržen pro selhání na paniku.
	panic("problém")

	// Běžné použití paniky je k přerušení, pokud funkce vrátí chybovou
	// hodnotu kterou nevíme jak (nebo nechceme) ošetřit.
	// Zde je příklad vyvolání `panic`
	// když dostaneme neočekávanou chybu během vytváření nového
	// souboru.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
