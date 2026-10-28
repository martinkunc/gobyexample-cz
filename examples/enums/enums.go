// _Výčtové typy_ (enums) jsou speciálním případem
// [součtových typů](https://en.wikipedia.org/wiki/Algebraic_data_type).
// Výčtový typ je typ, který má pevný počet možných
// hodnot, každou s odlišným názvem. Go nemá výčtový typ
// jako samostatnou vlastnost jazyka, ale výčtové typy
// se dají jednoduše implementovat pomocí existujících
// idiomů jazyka.

package main

import "fmt"

// Náš výčtový typ `ServerState` má podkladový typ `int`.
type ServerState int

// Možné hodnoty `ServerState` jsou definované jako
// konstanty. Speciální klíčové slovo [iota](https://go.dev/ref/spec#Iota)
// automaticky generuje po sobě jdoucí hodnoty konstant;
// v tomto případě 0, 1, 2 a tak dále.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// Díky implementaci interface [fmt.Stringer](https://pkg.go.dev/fmt#Stringer)
// je možné hodnoty `ServerState` vypsat nebo převést
// na řetězce.
//
// Pokud je možných hodnot hodně, může to být těžkopádné.
// V takových případech je možné použít
// [nástroj stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
// spolu s `go:generate` a celý proces automatizovat.
// Delší vysvětlení najdete v [tomto článku](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate).
var stateName = map[ServerState]string{
	StateIdle:      "nečinný",
	StateConnected: "připojený",
	StateError:     "chyba",
	StateRetrying:  "opakování",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// Pokud máme hodnotu typu `int`, nemůžeme ji předat
	// funkci `transition` - překladač si bude stěžovat na
	// nesoulad typů. To u výčtových typů poskytuje určitou
	// míru typové bezpečnosti už při překladu.

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition napodobuje přechod mezi stavy serveru;
// převezme existující stav a vrátí nový stav.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Předpokládejme, že tady kontrolujeme nějaké
		// podmínky, abychom určili další stav...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("neznámý stav: %s", s))
	}
}
