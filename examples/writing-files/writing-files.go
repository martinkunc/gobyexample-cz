// Zápis do souborů v Go sleduje podobné vzory jako
// jsme viděli dříve u čtení.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Pro začátek tady je jak vypsat řetězec (nebo prostě
	// bajty) do souboru.
	d1 := []byte("ahoj\ngo\n")
	path1 := filepath.Join(os.TempDir(), "dat1")
	err := os.WriteFile(path1, d1, 0644)
	check(err)

	// Tento trochu méně sofistikovaný zápis otevře soubor
	// pro zápis.
	path2 := filepath.Join(os.TempDir(), "dat2")
	f, err := os.Create(path2)
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
	check(err)
	fmt.Printf("zapsáno %d bajtů\n", n3)

	// Pro vyprázdnění zápisů na stabilní úložiště
	// zavolej `Sync`
	f.Sync()

	// `bufio` umožňuje bufferované Writery podobné bufferovaným
	//  Readerům které jsme viděli dříve.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("bufferováno\n")
	check(err)
	fmt.Printf("zapsáno %d bajtů\n", n4)

	// Použijte `Flush` pro zajištění že všechny bufferované operace
	// byly na vnitřním Writeru provedené.
	w.Flush()

}
