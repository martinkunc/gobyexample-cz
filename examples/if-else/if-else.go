// Větvení s příkazy `if` (jestliže) a `else` (jinak) je
// v Go přímočaré

package main

import "fmt"

func main() {

	// Tady je základní příklad
	if 7%2 == 0 {
		fmt.Println("7 je sudá")
	} else {
		fmt.Println("7 je lichá")
	}

	// Můžeme napsat `if` bez `else`
	if 8%4 == 0 {
		fmt.Println("8 je dělitelná 4")
	}

	// Podmínky může předcházet příkaz; proměnné
	// deklarované v tomto příkazu jsou dostupné ve všech
	// větvích
	if num := 9; num < 0 {
		fmt.Println(num, "je záporná")
	} else if num < 10 {
		fmt.Println(num, "obsahuje 1 číslici")
	} else {
		fmt.Println(num, "obsahuje víc číslic")
	}
}

// Povšimněte si, že v Go nejsou okolo podmínky nutné
// závorky, naproti tomu složené závorky jsou povinné.
