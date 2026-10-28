// Go podporuje _vkládání_ (embedding) struktur a interface
// pro plynulejší vyjádření _skládání_ (composition) typů.
// Nezaměňujte to s [`//go:embed`](embed-directive), což je
// direktiva Go zavedená ve verzi Go 1.16+ pro vkládání
// souborů a adresářů do binárky aplikace.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base s num=%v", b.num)
}

// `container` _vkládá_ `base`. Vložení vypadá jako
// položka bez jména.
type container struct {
	base
	str string
}

func main() {

	// Když vytváříme struktury pomocí literálů, musíme
	// vložení inicializovat explicitně; jako název
	// položky zde slouží vložený typ.
	co := container{
		base: base{
			num: 1,
		},
		str: "nějaké jméno",
	}

	// K položkám `base` můžeme přistupovat přímo na `co`,
	// např. `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Případně můžeme vypsat celou cestu pomocí názvu
	// vloženého typu.
	fmt.Println("také num:", co.base.num)

	// Protože `container` vkládá `base`, metody `base`
	// se stávají i metodami `container`u. Tady voláme
	// metodu, která byla vložena z `base`, přímo na `co`.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Vkládání struktur s metodami se dá použít pro
	// propůjčení implementací interface jiným strukturám.
	// Tady vidíme, že `container` teď implementuje
	// interface `describer`, protože vkládá `base`.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
