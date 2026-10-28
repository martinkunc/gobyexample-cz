// Go nabízí rozsáhlou podporu pro časy a doby trvání;
// zde jsou nějaké příklady.

package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	// Začneme získáním aktuálního času.
	now := time.Now()
	p(now)

	// Můžeš vytvořit strukturu `time` poskytnutím
	// roku, měsíce, dne, atd. Časy jsou vždy spojeny s
	// nějakou lokací (`Location`), např. časovou zónou.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Je možné získat různé komponenty z hodnoty time
	// očekávaným způsobem.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Také je dostupný den v týdnu Pondělí-Neděle (`weekday`).
	p(then.Weekday())

	// Tyto metody spolu porovnávají dva časy, testují jestli
	// se první stal před, po a nebo ve stejný čas
	// jako druhý.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Metody `Sub` vrací trvání (`Duration`) reprezentující
	// interval mezi dvěma časy.
	diff := now.Sub(then)
	p(diff)

	// Můžeme vypočítat délku trvání v různých jednotkách.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Můžeš použít `Add` pro to posun času dopředu o danou dobu,
	// nebo s `-` pro posun času o danou dobu zpět.
	p(then.Add(diff))
	p(then.Add(-diff))
}
