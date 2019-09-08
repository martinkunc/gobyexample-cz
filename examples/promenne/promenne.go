// V Go, _proměnné_ jsou explicitně deklarované a to je
// využíváné překladačem pro např. kontrolu typové
// správnosti u volání funkcí.

package main

import "fmt"

func main() {

	// `var` deklaruje 1 nebo více proměnných.
	var a = "počáteční"
	fmt.Println(a)

	// Je možné deklarovat více proměnných najednou.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go odvodí typ inicializavaných proměnných
	var d = true
	fmt.Println(d)

	// Proměnné deklarované bez svých počátečních stavů
	// mají _nulové-hodnoty_. Například,
	// nulová hodnota proměnné typu `int` je `0`.
	var e int
	fmt.Println(e)

	// Zápis `:=` je zkratkou pro deklaraci a inicializaci
	// proměnné, např. v tomto případě
	// `var f string = "apple"`
	f := "apple"
	fmt.Println(f)
}
