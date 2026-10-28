// Někdy naše Go programy potřebují spustit jiné
// procesy.

package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// Začneme jednoduchým příkazem který nemá žádné
	// argumenty nebo vstup a prostě vypisuje něco na
	// stdout. Pomocný příkaz `exec.Command` vytváří
	// objekt pro reprezentaci tohoto externího procesu.
	dateCmd := exec.Command("date")

	// Metoda `Output` spustí příkaz, počká na jeho
	// ukončení a posbírá jeho standardní výstup.
	// Pokud nedošlo k žádným chybám, bude `dateOut`
	// obsahovat bajty s informací o datu.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// `Output` a další metody `Command` vrátí
	// `*exec.Error`, pokud byl při spouštění příkazu
	// problém (např. špatná cesta), a `*exec.ExitError`,
	// pokud se příkaz spustil, ale skončil s nenulovým
	// návratovým kódem.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		if e, ok := errors.AsType[*exec.Error](err); ok {
			fmt.Println("spuštění selhalo:", e)
		} else if e, ok := errors.AsType[*exec.ExitError](err); ok {
			exitCode := e.ExitCode()
			fmt.Println("návratový kód příkazu =", exitCode)
		} else {
			panic(err)
		}
	}

	// Dál se podíváme na trochu komplikovanější případ
	// kde propojíme pomocí pipe data do externího procesu
	// na jeho `stdin` a posbíráme výsledky z jeho `stdout`.
	grepCmd := exec.Command("grep", "hello")

	// Tady explicitně vezmeme vstupní a výstupní pipe,
	// nastartujeme proces, zapíšeme do něho nějaký vstup,
	// načteme výsledný výstup, a na závěr počkáme na ukončení
	// tohoto procesu.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// V našem příkladu výše jsme vynechali kontroly chyb, ale
	// pro všechny můžete použít obvyklý vzor `if err != nil`.
	// Také sbíráme výsledky jen ze `StdoutPipe`,
	// ale mohli byste přesně stejným způsobem sbírat i
	// `StderrPipe`.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Všimněte si, že když spouštíme příkaz musíme
	// poskytnout explicitně definovaný příkaz a
	// pole argumentů v porovnání s prostým předáním
	// jednoho řetězce s příkazovou řádkou. Pokud chcete
	// spustit plný příkaz pomocí řetězce, můžete použít
	// volbu `-c` `bash` shellu:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
