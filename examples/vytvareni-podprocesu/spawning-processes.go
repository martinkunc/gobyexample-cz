// Někdy naše Go programu potřebují zplodit jiné, ne-Go
// procesy. Například, syntaktické podbarvení na téhle
// stránce je [implementovaný](https://github.com/mmcgrana/gobyexample/blob/master/tools/generate.go)
// pomocí spouštění [`pygmentize`](http://pygments.org/)
// podprocesu z Go programu. Pojďme se podívat na pár
// příkladů spouštění podprocesů v Go.

package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {

	// Začneme jednoduchým příkazem který nemá žádné
	// argumenty nebo vstup a prostě vypisuje něco na
	// stdout. Pomocný příkaz `exec.Command` vytváří
	// objekt pro reprezentaci tohoto externího procesu.
	dateCmd := exec.Command("date")

	// `.Output` je další z pomocných funkcí které ošetřují
	// běžný případ běhu příkazu s čekáním na jeho ukončení,
	// a sběrem jeho výstupu. Pokud nedošlo k žádným chybám,
	// bude `dateOut` obsahovat bajty s informací o datumu.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Dál se podíváme na trochu kompikovanější případ
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
	grepBytes, _ := ioutil.ReadAll(grepOut)
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
