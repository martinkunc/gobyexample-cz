// _range_ iteruje přes elementy v různých datových
// strukturách. Pojďmě se podívat jako použít `range`
// s některými datovými strukturami o kterých jsme se už
// dozvěděli.

package main

import "fmt"

func main() {

	// Tady používáme `range` pro součet čísel ve výřezu.
	// Jde to použít stejně i u polí.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("součet:", sum)

	// `range` na polích a výřezech poskytuje oboje, index
	// a hodnotu pro každý záznam. Výše jsme index
	// nepotřebovali, tak jsme ho ignorovali s použitím
	// zahazovacího identifikátoru `_`. Přesto se někdy
	// index může hodit.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` v mapách iteruje přes dvojice klíč/hodnota.
	kvs := map[string]string{"a": "ananas", "b": "banán"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` může také iterovat jen přes klíče v mapě.
	for k := range kvs {
		fmt.Println("klíč:", k)
	}

	// `range` na řetězcích iteruje přes Unicode kódové
	// body. První hodnota je index počátečního byte
	// `runy` a druhá je samotná `runa`.
	// *Pozn: Unicode kódový bod je pozice v tabulce kódu,
	// může reprezentovat např. znak. Runa je Go typ pro
	// jeden codepoint reprezentovaný jako znak.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
