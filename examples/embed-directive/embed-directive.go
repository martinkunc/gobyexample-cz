// `//go:embed` je [direktiva
// překladače](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives), která
// programům umožňuje při sestavení zahrnout do Go binárky libovolné soubory
// a adresáře. Více o direktivě embed si můžete přečíst
// [zde](https://pkg.go.dev/embed).
package main

// Importujeme balíček `embed`; pokud z tohoto balíčku nepoužíváte žádné
// exportované identifikátory, můžete použít prázdný import `_ "embed"`.
import (
	"embed"
)

// Direktivy `embed` přijímají cesty relativní k adresáři obsahujícímu
// zdrojový soubor Go. Tato direktiva vloží obsah souboru do proměnné
// typu `string`, která následuje hned za ní.
//
//go:embed folder/single_file.txt
var fileString string

// Nebo vložíme obsah souboru do `[]byte`.
//
//go:embed folder/single_file.txt
var fileByte []byte

// Pomocí zástupných znaků můžeme vložit i více souborů nebo dokonce
// adresáře. K tomu slouží proměnná typu [embed.FS](https://pkg.go.dev/embed#FS),
// která implementuje jednoduchý virtuální souborový systém.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// Vypíšeme obsah `single_file.txt`.
	print(fileString)
	print(string(fileByte))

	// Načteme několik souborů z vloženého adresáře.
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
