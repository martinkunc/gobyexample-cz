// Go poskytuje zabudovanou podporu pro [base64
// kódování/dekódování](http://en.wikipedia.org/wiki/Base64).

package main

// Tato syntaxe importuje balíček `encoding/base64` se jménem
// `b64` místo výchozího `base64`. Ušetří nám to níže nějaké
// místo.
import b64 "encoding/base64"
import "fmt"

func main() {

	// Tady je `string` který budeme kódovat/dekódovat.
	data := "abc123!?$*&()'-=@~"

	// Go podporuje standardní i URL-kompatibilní
	// base64. Tady je jak zakódovat s použitím
	// standardního enkóderu. Enkóder vyžaduje `[]byte`
	// takže převedeme náš `string` na tento typ.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Dekódování může vrátit chybu, kterou můžeme zkontrolovat,
	// pokud už nevíš že je vstup dobře naformátovaný.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// Toto kóduje/dekóduje s použitím URL-kompatibilního base64
	// formátu.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
