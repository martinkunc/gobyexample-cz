// _Interface_ (rozhraní) jsou pojmenované kolekce
// signatur metod.

package main

import "fmt"
import "math"

// Zde je základní interface pro geometrické útvary.
type geometrie interface {
	plocha() float64
	obvod() float64
}

// V našem příkladu implementujeme tento interface pro
// typy `obdelnik` and `kruznice` types.
type obdelnik struct {
	sirka, vyska float64
}
type kruznice struct {
	polomer float64
}

// Pro implementaci interface v Go potřebujeme jenom
// implementovat všechny metody interface.
// Tady implementujeme `geometrie` na `obdelnik`u.
func (r obdelnik) plocha() float64 {
	return r.sirka * r.vyska
}
func (r obdelnik) obvod() float64 {
	return 2*r.sirka + 2*r.vyska
}

// Implementace pro `kruznice`.
func (c kruznice) plocha() float64 {
	return math.Pi * c.polomer * c.polomer
}
func (c kruznice) obvod() float64 {
	return 2 * math.Pi * c.polomer
}

// Pokud má proměnná typ interface, potom na ni můžeme
// volat metody které jsou ve zmíněném interface. Zde je
// obecná funkce `zmer` která toho využívá pro práci s
// libovolnou `geometrie`.
func zmer(g geometrie) {
	fmt.Println(g)
	fmt.Println(g.plocha())
	fmt.Println(g.obvod())
}

func main() {
	r := obdelnik{sirka: 3, vyska: 4}
	c := kruznice{polomer: 5}

	// Oba strukturové typy `kruznice` i `obdelnik`
	// implementují interface `geometrie` takže můžeme
	// použít instance těchto struktur jako argumenty
	// funkci `zmer`.
	zmer(r)
	zmer(c)
}
