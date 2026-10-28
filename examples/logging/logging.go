// Standardní knihovna Go poskytuje přímočaré nástroje
// pro výpis logů z Go programů: balíček
// [log](https://pkg.go.dev/log) pro volně formátovaný
// výstup a balíček [log/slog](https://pkg.go.dev/log/slog)
// pro strukturovaný výstup.
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	// Prosté zavolání funkcí jako `Println` z balíčku
	// `log` použije _standardní_ logger, který je už
	// předem nastavený pro rozumný výstup logů do
	// `os.Stderr`. Další metody jako `Fatal*` nebo
	// `Panic*` po zalogování program ukončí.
	log.Println("standardní logger")

	// Loggery je možné nastavit pomocí _příznaků_
	// (flags), které určují formát jejich výstupu.
	// Ve výchozím stavu má standardní logger nastavené
	// příznaky `log.Ldate` a `log.Ltime`, které jsou
	// shrnuté v `log.LstdFlags`. Můžeme mu například
	// změnit příznaky tak, aby vypisoval čas
	// s přesností na mikrosekundy.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("s mikrosekundami")

	// Podporuje také výpis názvu souboru a řádku,
	// ze kterého je funkce `log` volána.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("se souborem/řádkem")

	// Může se hodit vytvořit si vlastní logger a předávat
	// ho dál. Při vytváření nového loggeru můžeme nastavit
	// _prefix_, který odliší jeho výstup od ostatních
	// loggerů.
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("z mylog")

	// Prefix můžeme nastavit i u existujících loggerů
	// (včetně standardního) pomocí metody `SetPrefix`.
	mylog.SetPrefix("ohmy:")
	mylog.Println("z mylog")

	// Loggery mohou mít vlastní výstupní cíle;
	// funguje jakýkoli `io.Writer`.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// Toto volání zapíše výstup logu do `buf`.
	buflog.Println("ahoj")

	// A tohle ho skutečně zobrazí na standardním výstupu.
	fmt.Print("z buflog:", buf.String())

	// Balíček `slog` poskytuje _strukturovaný_ výstup
	// logů. Například logování ve formátu JSON je
	// přímočaré.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("ahoj")

	// Kromě zprávy může výstup `slog` obsahovat
	// libovolný počet dvojic klíč=hodnota.
	myslog.Info("znovu ahoj", "key", "val", "age", 25)
}
