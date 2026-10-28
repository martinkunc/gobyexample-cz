// Během běhu programu často chceme vytvářet data, která
// po skončení programu už nejsou potřeba. K tomu se
// hodí *dočasné soubory a adresáře*, protože časem
// nezaneřádí souborový systém.

package main

import (
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

	// Nejjednodušší způsob, jak vytvořit dočasný soubor,
	// je zavolat `os.CreateTemp`. Ten soubor vytvoří *a*
	// otevře ho pro čtení i zápis. Jako první argument
	// předáme `""`, takže `os.CreateTemp` vytvoří soubor
	// ve výchozím umístění pro náš operační systém.
	f, err := os.CreateTemp("", "sample")
	check(err)

	// Vypíšeme název dočasného souboru. Na operačních
	// systémech založených na Unixu to bude
	// pravděpodobně adresář `/tmp`. Název souboru začíná
	// prefixem předaným jako druhý argument
	// `os.CreateTemp` a zbytek se zvolí automaticky tak,
	// aby souběžná volání vždy vytvořila různé názvy.
	fmt.Println("Název dočasného souboru:", f.Name())

	// Po skončení soubor uklidíme. Operační systém
	// dočasné soubory pravděpodobně po nějaké době uklidí
	// sám, ale je dobrým zvykem to udělat explicitně.
	defer os.Remove(f.Name())

	// Do souboru můžeme zapsat nějaká data.
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// Pokud chceme zapisovat hodně dočasných souborů,
	// možná dáme přednost vytvoření dočasného *adresáře*.
	// Argumenty `os.MkdirTemp` jsou stejné jako
	// u `CreateTemp`, ale vrací *název* adresáře místo
	// otevřeného souboru.
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Název dočasného adresáře:", dname)

	defer os.RemoveAll(dname)

	// Teď můžeme vytvářet názvy dočasných souborů tak,
	// že před ně přidáme náš dočasný adresář.
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
