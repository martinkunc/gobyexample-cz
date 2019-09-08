// Často chceme v Go provádět kód v nějakám čase v
// budoucnu, a nebo opakovaně v nějakém intervalu.
// Dva zabudované typy v Go, časovače _timers_ a opakující
// časovače _tickers_ takové případy usnadňují.
// Nejdřív se podíváme na časovače (timers) a potom
// na [opakující časovače (tickers)](opakujici-casovace-tickers).

package main

import "time"
import "fmt"

func main() {

	// Časovače představují jednu událost v budoucnu.
	// Sdělíte časovaci jak dlouho chcete čekat, a on vám
	// poskytne kanál kterému bude po tom čase bude
	// posláno oznámení.
	// Tento časovač bude čekat 2 vteřiny.
	timer1 := time.NewTimer(2 * time.Second)

	// Výraz `<-timer1.C` bude blokovat na kanálu `C`
	// časovače, dokud ten nepošle hodnotu oznamující
	// že čas vypršel.
	<-timer1.C
	fmt.Println("Timer 1 vypršel")

	// Jestli jste jen chtěli počkat, tak jste mohli
	// použít `time.Sleep`. Jeden důvod kde časovač může
	// být užitečný je že časovač můžete zrušit ještě
	// předtím než čas vyprší.
	// Příklad toho je tady.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 vypršel")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 byl zastaven")
	}
}
