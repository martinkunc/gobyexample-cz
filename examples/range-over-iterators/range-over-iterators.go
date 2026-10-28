// Od verze 1.23 přidalo Go podporu pro
// [iterátory](https://go.dev/blog/range-functions),
// díky kterým můžeme pomocí `range` iterovat prakticky
// přes cokoliv!

package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

// Podívejme se znovu na typ `List` z
// [předchozího příkladu](generics). V tom příkladu jsme
// měli metodu `AllElements`, která vracela výřez všech
// prvků v seznamu. S iterátory v Go to můžeme udělat
// lépe - jak je ukázáno níže.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All vrací _iterátor_, což je v Go funkce se
// [speciální signaturou](https://pkg.go.dev/iter#Seq).
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// Funkce iterátoru přebírá jako parametr jinou
		// funkci, podle konvence nazvanou `yield` (název
		// ale může být libovolný). Zavolá `yield` pro každý
		// prvek, přes který chceme iterovat, a sleduje
		// návratovou hodnotu `yield` kvůli případnému
		// předčasnému ukončení.
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Iterace nevyžaduje podkladovou datovou strukturu
// a nemusí být dokonce ani konečná! Tady je funkce
// vracející iterátor přes Fibonacciho čísla: běží tak
// dlouho, dokud `yield` vrací `true`.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Protože `List.All` vrací iterátor, můžeme ho použít
	// v běžném cyklu `range`.
	for e := range lst.All() {
		fmt.Println(e)
	}

	// Balíčky jako [slices](https://pkg.go.dev/slices) mají
	// řadu užitečných funkcí pro práci s iterátory.
	// Například `Collect` převezme libovolný iterátor
	// a posbírá všechny jeho hodnoty do výřezu.
	all := slices.Collect(lst.All())
	fmt.Println("vše:", all)

	// Balíčky standardní knihovny teď také nabízejí
	// pomocné funkce s iterátory. Například
	// `strings.SplitSeq` iteruje přes části řetězce, aniž
	// by nejdřív vytvářel výřez s výsledkem.
	for part := range strings.SplitSeq("go-by-example", "-") {
		fmt.Printf("část: %s\n", part)
	}

	for n := range genFib() {

		// Jakmile cyklus narazí na `break` nebo předčasný
		// `return`, funkce `yield` předaná iterátoru
		// vrátí `false`.
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
