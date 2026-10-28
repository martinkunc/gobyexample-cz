// Zápis do souborů v Go sleduje podobné vzory jako
// jsme viděli dříve u čtení.

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Pro začátek tady je jako vypsat řetězec, nebo prostě
	// bajty) do souboru.
	d1 := []byte("ahoj\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// Tento trochu méně sofistikovaný zápis otevře soubor
	// pro zápis.
	f, err := os.Create("/tmp/dat2")
	check(err)

	// V Go je idiomatické odložit `Close` hned po
	// otevření souboru.
	defer f.Close()

	// Jak byste asi čekali, je možné zapsat s `Write` výřez
	// bajtů.
	d2 := []byte{110, 196, 155, 106, 97, 107, 195,
		169, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("zapsáno %d bajtů\n", n2)

	// Také existuje `WriteString`.
	n3, err := f.WriteString("writes\n")
	fmt.Printf("zapsáno %d bajtů\n", n3)

	// Pro vyprázdnění zápisů na stabilní úložiště
	// zavolej `Sync`
	f.Sync()

	// `bufio` umožňuje bufferované Writery podobné bufferovaným
	//  Readerům které jsme viděli dříve.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("bufferováno\n")
	fmt.Printf("zapsáno %d bajtů\n", n4)

	// Použijte `Flush` pro zajištění že všechny bufferované operace
	// byly na vnitřním Writeru provedené.
	w.Flush()

}
