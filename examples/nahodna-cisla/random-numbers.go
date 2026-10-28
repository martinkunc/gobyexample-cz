// Balíček `math/rand` v Go poskytuje
// generování [pseudonáhodných čísel](http://en.wikipedia.org/wiki/Pseudorandom_number_generator)

package main

import "time"
import "fmt"
import "math/rand"

func main() {

	// Například, `rand.Intn` vrátí náhodné n typu `int`,
	// tak že platí `0 <= n < 100`.
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// `rand.Float64` vrátí `float64` `f`,
	// tak že platí `0.0 <= f < 1.0`.
	fmt.Println(rand.Float64())

	// Tohle se dá použít pro generování náhodných floatů v
	// jiném rozsahu, například `5.0 <= f' < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Výchozí číselný generátor je deterministický,
	// takže ve výchozím stavu vyprodukuje pokaždé stejnou
	// sekvenci čísel. Pro tvorbu měnících se sekvencí,
	// poskytněte generátoru semínko které se mění.
	// Povšimněte se že není bezpečné tento generátor používat pro
	// náhodná čísla která by měla být tajná.
	// Pro ty použij `crypto/rand`.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Zavolej výslednou `rand.Rand` stejně jako
	// funkce v balíčku `rand`.
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// Pokud inicializuješ zdroj se stejným číslem, produkuje
	// stejnou sekvenci náhodných čísel.
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
