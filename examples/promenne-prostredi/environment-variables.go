// [Proměnné prostředí](http://en.wikipedia.org/wiki/Environment_variable)
// jsou univerzální mechanismus pro [předávání informací o konfiguraci
// Unixovým programům](http://www.12factor.net/config).
// Pojďme se podívat jak nastavit, získat a vypsat proměnné prostředí.

package main

import "os"
import "strings"
import "fmt"

func main() {

	// Pro nastavení dvojice klíč/hodnota, použijte `os.Setenv`. Pro
	// získání hodnoty pro klíč, použijte `os.Getenv`. Ta vrátí
	// prázdný řetězec pokud nebude klíč v prostředí existovat.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// Použijte `os.Environ` pro vypsání všech dvojic key/value v
	// prostředí. Toto vrátí výřez řetězců ve formě `KLÍČ=hodnota`.
	// Můžeš na ně použít `strings.Split` pro získání
	// klíče a hodnoty. Tady vypíšeme všechny klíče.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
