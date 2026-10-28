// V předchozím příkladu jsme se dívali na
// [vytváření externích podprocesů](vytvareni-podprocesu).
// To děláme když potřebujeme externí proces přístupný z
// běžícího Go procesu. Někdy prostě chceme kompletně
// nahradit aktuální Go proces jiným (možné ne-Go).
// procesem. Abychom to udělali použijeme Go
// implementaci klasické funkce
// <a href="http://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>.

package main

import "syscall"
import "os"
import "os/exec"

func main() {

	// Pro náš příklad použijeme exec `ls`. Go vyžaduje absolutní
	// cestu k binárnímu programu který chcete spustit, takže
	// použijeme `exec.LookPath` pro jejího nalezení (pravděpodobně
	// `/bin/ls`).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` vyžaduje argumentz ve formě výřezu (a ne jako jeden
	// velký řetězec). Zkusíme dát `ls` pár běžných argumentů.
	// Všimněte si, že první argument by měl být název
	// programu.
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` jaké potřebuje předat množinu [proměnných prostředí](promenne-prostredi).
	// Zde prostě poskytneme naše aktuální prostředí.
	env := os.Environ()

	// Zde je vlastní volání `syscall.Exec`. Pokud je toto volání
	// úspěšné, běh našeho procesu zde končí
	// a bude nahrazený procesem `/bin/ls -a -l -h`
	// Když došlo k chybě, dostaneme návratovou hodnotu.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
