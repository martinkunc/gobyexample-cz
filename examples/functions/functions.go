// _Funkce_ jsou v Go to hlavní. Ukážeme si je na
// pár různých příkladech.

package main

import "fmt"

// Tady je funkce která bere dva `int`egery a vrací
// jejich součet jako `int`.
func plus(a int, b int) int {

	// Go vyžaduje explicitní return, tj. nevrátí
	// automaticky hodnotu posledního výrazu.
	return a + b
}

// Pokud máte víc po sobě jdoucích parametrů stejného
// typu, můžete vynechat název typu pro parametry stejných
// typů až na poslední, který pak určuje typ.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Pro volání funkce použijte přesně jak byste čekali:
	// `jméno(parametry)`.
	vysl := plus(1, 2)
	fmt.Println("1+2 =", vysl)

	vysl = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", vysl)
}
