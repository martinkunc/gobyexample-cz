// Go podporuje <em><a href="http://en.wikipedia.org/wiki/Pointer_(computer_programming)">ukazatele</a></em>,
// díky nim je možné předávat si v programu odkazy na
// hodnoty a záznamy.

package main

import "fmt"

// Jak ukazatelé fungují v protikladu k proměnným si
// ukážeme na příkladu dvou funkcí: `nulujhodn` a
// `nulujukaz`. `nulujhodn` má
// parametr `int` a argument ji tak bude předaný jako
// hodnota. `nulujhodn` dostane kopii `ival` odlišnou od
// té ve volající funkci
func nulujhodn(ival int) {
	ival = 0
}

// Naproti tomu `nulujukaz` má parametr typu `*int`,
// značící že přijímá ukazatel na `int`. V těle funkce
// potom kód `*iptr` _dereferencuje_ ukazatel z jeho
// adresy v paměti na aktuální hodnotu na této adrese.
// Přiřazením hodnoty do dereferencovaného ukazatele
// změní hodnotu na odkazované adrese.
func nulujukaz(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("počáteční:", i)

	nulujhodn(i)
	fmt.Println("nulujhodn:", i)

	// Zápis `&i` vrací adresu proměnné `i` v paměti,
	// např. ukazatel na `i`.
	nulujukaz(&i)
	fmt.Println("nulujukaz:", i)

	// Pointers can be printed too.
	fmt.Println("ukazatel:", &i)
}
