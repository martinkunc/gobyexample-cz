// Od verze 1.18 přidalo Go podporu pro _generika_
// (generics), známá také jako _typové parametry_.

package main

import "fmt"

// Jako příklad generické funkce, `SlicesIndex` převezme
// výřez libovolného typu `comparable` a prvek tohoto
// typu a vrátí index prvního výskytu v v s, nebo -1,
// pokud tam není. Omezení `comparable` znamená, že
// hodnoty tohoto typu můžeme porovnávat operátory `==`
// a `!=`. Podrobnější vysvětlení této typové signatury
// najdete v [tomto blog postu](https://go.dev/blog/deconstructing-type-parameters).
// Všimněte si, že tato funkce ve standardní knihovně
// existuje jako [slices.Index](https://pkg.go.dev/slices#Index).
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// Jako příklad generického typu, `List` je jednosměrně
// vázaný seznam s hodnotami libovolného typu.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Na generických typech můžeme definovat metody stejně
// jako na běžných typech, ale typové parametry musíme
// ponechat na místě. Typ je `List[T]`, ne `List`.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// AllElements vrátí všechny prvky seznamu `List` jako
// výřez. V dalším příkladu uvidíme idiomatičtější způsob,
// jak iterovat přes všechny prvky uživatelských typů.
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// Při volání generických funkcí se často můžeme
	// spolehnout na _odvození typů_ (type inference).
	// Všimněte si, že při volání `SlicesIndex` nemusíme
	// uvádět typy pro `S` a `E` - překladač je odvodí
	// automaticky.
	fmt.Println("index zoo:", SlicesIndex(s, "zoo"))

	// ... i když bychom je mohli uvést i explicitně.
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("seznam:", lst.AllElements())
}
