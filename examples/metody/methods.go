// Go podporuje _metody_ definované na typech struktur.

package main

import "fmt"

type obdel struct {
	sirka, vyska int
}

// Tato metoda `plocha` má typ příjemce `*obdel`.
func (r *obdel) plocha() int {
	return r.sirka * r.vyska
}

// Metody mohou být definovane jako příjemce pro typy buď
// ukazatel, nebo hodnota. Zde je příklad hodnotového
// příjemce.
func (r obdel) obvod() int {
	return 2*r.sirka + 2*r.vyska
}

func main() {
	r := obdel{sirka: 10, vyska: 5}

	// Tady voláme obě dvě metody definované pro naší
	// strukturu.
	fmt.Println("plocha: ", r.plocha())
	fmt.Println("obvod:", r.obvod())

	// Go se automaticky stará o konverzi mezi voláním
	// metod definovaných jako hodnotoví nebo ukazateloví
	// příjemci. Metoda na ukazatelovém příjemci se vám
	// může hodit abyste se vyhli kopírování struktury při
	// volání a nebo abyste dovolili metodě měnit
	// příjímající strukturu.
	rp := &r
	fmt.Println("plocha: ", rp.plocha())
	fmt.Println("obvod:", rp.obvod())
}
