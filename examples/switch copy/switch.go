// _Příkazy switch_ vyjadřují podmínky s mnoha větvemi.

package main

import "fmt"
import "time"

func main() {

	// Tady je základní `switch`.
	i := 2
	fmt.Print("Píšu ", i, " jako ")
	switch i {
	case 1:
		fmt.Println("jedna")
	case 2:
		fmt.Println("dvě")
	case 3:
		fmt.Println("tři")
	}

	// Můžete použít čárku pro oddělení více výrazů ve
	// stejném větvi `case`. V tomto příkladě používáme
	// také volitelnou case větev `default`.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Je víkend")
	default:
		fmt.Println("Je pracovní den")
	}

	// Příkaz `switch` bez výrazu je alternativní způsob
	// jak vyjádřit logiku if/else. Také tady ukazujeme
	// jak výrazy `case` mohou být ne-konstantní.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Je před polednem")
	default:
		fmt.Println("Je po poledni")
	}

	// Typový příkaz `switch` porovnává typy místo hodnot.
	// Můžete jej použit pro zjištění  typu hodnoty
	// interface.  V tomto příkladu, proměnná `t` bude mít
	// typ podle své odpovídající klauzule
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Jsem boolean")
		case int:
			fmt.Println("Jsem int")
		default:
			fmt.Printf("Neznám typ %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hej")
}
