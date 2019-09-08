// Go podporuje formátování a parsování času pomocí
// formátu rozložení založeném na vzorcích.

package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	// Tady je základní příklad formátování času
	// podle RFC3339, s využitím odpovídající konstanty
	// formátu rozložení.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// Parsování času používá stejné hodnoty rozložení jako `Format`.
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format` i `Parse` používají na příkladech založená rozložení.
	// Obvykle použiješ konstantu z `time` pro tato rozložení, ale
	// také můžeš použít vlastní rozložení. Rozložení musí používat
	// referenční čas `Mon Jan 2 15:04:05 MST 2006` pro definování
	// vzorku s kterým by se měl formátovat/parsovat daný čas/řetězec.
	// Čas příkladu musí být přesně jak je ukázáno: rok 2006,
	// 15 pro hodinu, Monday pro den v týdnu, atd.
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// Pro čistě číselnou reprezentaci také můžeš použít
	// standardní formátování řezězců s extrahovanými
	// komponentami hodnoty času.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// `Parse` vrátí chybu při poškozeném vstupu
	// vysvětlující problém s parsováním.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
