// Balíček `filepath` poskytuje funkce pro parsování
// a vytváření *cest k souborům* způsobem který je přenosný
// mezi operačními systémy; například `dir/file` na Linuxu
// vs. `dir\file` na Windows.
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// Pro vytváření cest portovatelným způsobem by měl
	// být použítý `Join`. Ten převezme libovolný počet
	// parametrů a vztvoří z nich hierarchickou cestu.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Místo ručního spojování `/`s nebo `\`s bys měl vždycky
	// použít `Join`. Kromě portability, `Join` navíc také
	// cestu normalizuje a odstraní nadbytečné oddělovace a
	// změny adresářů.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` a `Base` se dají použít pro rozdělení cesty
	// na adresář a soubor. Alternativně `Split` vrátí
	// obojí v jednom volání.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// Můžeme zkontrolovat jestli je cesta absolutní.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Některé názvy souborů mají po tečce i příponu. Od
	// takových jmen můžeme příponu oddělit pomocí `Ext`.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Pro vrácení názvu souboru bez přípony,
	// použijeme `strings.TrimSuffix`.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` najde relativní cestu mezi *základem* a
	// *cílem*. Když není možné udělat relativní odkaz
	// ze základu na cíl, funkce vrátí chybu.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
