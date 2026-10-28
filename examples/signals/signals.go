// Někdy bychom chtěli, aby naše Go programy inteligentně
// ošetřovaly [Unixové signály](https://en.wikipedia.org/wiki/Unix_signal).
// Například můžeme chtít, aby server elegantně
// skončil, když dostane signál `SIGTERM`, nebo aby nástroj
// příkazové řádky přestal zpracovávat vstup, pokud obdrží `SIGINT`.
// Zde je moderní způsob, jak ošetřit signály pomocí kontextů.

package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	// `signal.NotifyContext` vrací kontext, který je zrušen
	// ve chvíli, kdy přijde jeden z uvedených signálů.
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Program zde bude čekat, dokud nedostane jeden
	// z nastavených signálů.
	fmt.Println("awaiting signal")
	<-ctx.Done()

	// `context.Cause` sděluje, proč byl kontext zrušen.
	// U zrušení vyvolaného signálem obsahuje i hodnotu
	// signálu.
	fmt.Println()
	fmt.Println(context.Cause(ctx))
	fmt.Println("exiting")
}
