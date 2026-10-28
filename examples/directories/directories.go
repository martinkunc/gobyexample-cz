// Go má několik užitečných funkcí pro práci
// s *adresáři* v souborovém systému.

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Vytvoříme nový podadresář v aktuálním pracovním
	// adresáři.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// Při vytváření dočasných adresářů je dobrým zvykem
	// odložit jejich odstranění pomocí `defer`.
	// `os.RemoveAll` smaže celý strom adresářů (podobně
	// jako `rm -rf`).
	defer os.RemoveAll("subdir")

	// Pomocná funkce pro vytvoření nového prázdného souboru.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// Pomocí `MkdirAll` můžeme vytvořit hierarchii
	// adresářů včetně rodičovských. Je to podobné
	// příkazu `mkdir -p` na příkazové řádce.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` vypíše obsah adresáře a vrátí výřez
	// objektů `os.DirEntry`.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Výpis subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` nám umožní změnit aktuální pracovní adresář,
	// podobně jako `cd`.
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Při výpisu *aktuálního* adresáře teď uvidíme obsah
	// `subdir/parent/child`.
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Výpis subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Pomocí `cd` se vrátíme tam, kde jsme začali.
	err = os.Chdir("../../..")
	check(err)

	// Adresář můžeme také procházet *rekurzivně*, včetně
	// všech jeho podadresářů. `WalkDir` přijímá callback
	// funkci, která zpracuje každý navštívený soubor nebo
	// adresář.
	fmt.Println("Procházení subdir")
	err = filepath.WalkDir("subdir", visit)
	check(err)
}

// `visit` se zavolá pro každý soubor nebo adresář,
// který `filepath.WalkDir` rekurzivně najde.
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
