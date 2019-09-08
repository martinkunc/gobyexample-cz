// [_Parametry příkazové řádky_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// jsou běžný způsob jak parametrizovat spouštění programů.
// Například, `go run hello.go` používá argumenty `run` a
// `hello.go` na program `go`.

package main

import "os"
import "fmt"

func main() {

	// `os.Args` zpřístupňuje přímý přístup k parametrům
	// příkazové řádky. Všimněte si, že první hodnotou v
	// tomto výřezu je cesta k programu a `os.Args[1:]`
	// obsahuje argumentu pro program.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// Můžete získat individuální argumenty normálním indexováním.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
