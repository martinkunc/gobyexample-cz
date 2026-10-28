// Go podporuje [_anonymni funkce_](http://en.wikipedia.org/wiki/Anonymous_function),
// které mohou výtvářet <a href="http://en.wikipedia.org/wiki/Closure_(computer_science)"><em>uzávěry</em></a>.
// Anonymní funkce jsou užitečné když chcete
// definovat funkci uvnitř funkce bez nutnosti ji
// pojmenovávat.

package main

import "fmt"

// Tato funkce `intSekv` vrací jinou funkci, kterou jsme
// definovali jako anonymní uvnitř funkce `intSekv`.
// Vracená funkce _uzavírá přes_ proměnnou `i` a vytváří
// tak uzávěru. *Pozn: Jinými slovy, hodnota uzavřené
// proměnné se zachovává přes další volání vrácené
// anonymní funkce.
func intSekv() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// Voláme `intSekv`, přiřazující výsledek (funkci)
	// do `dalsiInt`. Tato funkce zachytí svou vlastní
	// hodnotu `i`, která se bude aktualizovat pokaždé
	// během volání `dalsiInt`.
	dalsiInt := intSekv()

	// Všimněte si efektu uzávěry během několika volání
	// `dalsiInt`.
	fmt.Println(dalsiInt())
	fmt.Println(dalsiInt())
	fmt.Println(dalsiInt())

	// Pro ověření že stav je unikátní k té které funkci,
	// vytvoříme a otestujeme jednu novou.
	noveInty := intSekv()
	fmt.Println(noveInty())
}
