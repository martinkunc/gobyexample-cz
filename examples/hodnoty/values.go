// Go má různé hodnotové typy (value types) zahrnující
// řetězce (strings), celá čísla (integers), čísla s
// plovoucí des.čáskou (floats), pravdivostní hodnoty
// (booleans), atd. Tady je pár základních příkladů.

package main

import "fmt"

func main() {

	// Řetězce, které je možné spojovat pomocí `+`.
	fmt.Println("go" + "lang")

	// Integery a Floaty.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleany i s boolean operátory, jak byste čekali.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
