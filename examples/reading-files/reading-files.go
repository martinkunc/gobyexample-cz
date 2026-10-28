// Čtení z a zapisování do souborů jsou základní součásti
// mnoha Go programů. Nejprve se podíváme na nějaké
// příklady jak ze souborů číst.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Čtení souborů vyžaduje kontrolu chyb u většiny volání.
// Tato pomocná funkce nám zjednoduší naše kontroly chyb
// níže.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Snad nejzákladnější operací při čtení souborů je
	// nasátí kompletního obsahu souboru do paměti.
	path := filepath.Join(os.TempDir(), "dat")
	dat, err := os.ReadFile(path)
	check(err)
	fmt.Print(string(dat))

	// Často budete chtít více kontroly jak a jaké části
	// souboru jsou načítané. Pro tyto úkoly začněte s
	// otevřením souboru pomocí `Open` pro získání hodnoty
	// `os.File`.
	f, err := os.Open(path)
	check(err)

	// Načteme pár bytů ze začátku souboru.
	// Zkusí načíst až 4 a vrátí kolik jich bylo
	// ve skutečnosti načteno.
	b1 := make([]byte, 4)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d byty: %s\n", n1, string(b1[:n1]))

	// Můžete také přeskočit na danou pozici pomocí `Seek`
	// a pomocí `Read` načíst odtamtud.
	o2, err := f.Seek(5, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d byty @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// Další způsoby přeskakování jsou relativní k aktuální
	// pozici kurzoru,
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	// a relativní ke konci souboru.
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	// Balíček `io` poskytuje nějaké funkce které mohou
	// být užitečné pro čtení souborů. Například čtení
	// jako výše může být robustněji implementované pomocí
	// `ReadAtLeast`.
	o3, err := f.Seek(5, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d byty @ %d: %s\n", n3, o3, string(b3))

	// Neexistuje žádné zabudované přetočení. Dá se toho
	// dosáhnout pomocí skoku na pozici užitím
	// `Seek(0, io.SeekStart)` které je relativní k počátku
	// souboru.
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// Balíček `bufio` implementuje čtení pomocí readeru s
	// vyrovnávací pamětí které může být užitečné pro svou
	// efektivitu  při mnoha malých čteních ale pri kvůli
	// spoustě metod pro čtení které poskytuje.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(4)
	check(err)
	fmt.Printf("4 byty: %s\n", string(b4))

	// Zavře soubor když jsme hotovi. (Obvykle to bude
	// naplánováno hned po otevření s využitím `defer`).
	f.Close()
}
