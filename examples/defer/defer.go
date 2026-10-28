// _Defer_ se používá k zajištění že volání funkce je
// provedené později během programu, obvykle pro účely
// uvolňování. `defer` se často používá kde by se v jiných
// jazycích použilo např. `ensure` a `finally`.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Předpokládejme že jsme chtěli vytvořit soubor, zapsat
// do něho a potom když jsme hotovi ho zavřít. Tady je jak
// bychom to mohli udělat s `defer`.
func main() {

	// Ihned po získání objektu souboru od
	// `createFile`, pomocí `defer` odložíme uzavření tohoto
	// souboru zavoláním `closeFile`. To se provede při ukončení
	// volající funkce (`main`), potom co skončilo
	// `writeFile`.
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("vytváření")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("zápis")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("zavírání")
	err := f.Close()
	// Při zavírání souboru je důležité kontrolovat chyby,
	// a to i v odložené funkci.
	if err != nil {
		panic(err)
	}
}
