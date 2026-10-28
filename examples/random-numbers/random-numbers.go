// Balíček `math/rand/v2` v Go poskytuje
// generování [pseudonáhodných čísel](https://en.wikipedia.org/wiki/Pseudorandom_number_generator).

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// Například, `rand.IntN` vrátí náhodné n typu `int`,
	// tak že platí `0 <= n < 100`.
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// `rand.Float64` vrátí `float64` `f`,
	// tak že platí `0.0 <= f < 1.0`.
	fmt.Println(rand.Float64())

	// Tohle se dá použít pro generování náhodných floatů v
	// jiném rozsahu, například `5.0 <= f' < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Pokud chcete známé semínko, vytvořte nový
	// `rand.Source` a předejte ho konstruktoru `New`.
	// `NewPCG` vytvoří nový zdroj
	// [PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator),
	// který vyžaduje semínko ze dvou čísel typu `uint64`.
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
