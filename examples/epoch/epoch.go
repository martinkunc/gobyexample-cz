// Běžný požadavek v programech je získání počtu vteřin,
// milisekund, nebo nanosekund z
// [Unixového času](http://en.wikipedia.org/wiki/Unix_time).
// Tady je jako to udělat v Go.

package main

import "fmt"
import "time"

func main() {

	// Pro získání uplynulého času od Unix epochy v
	// sekundách nebo nanosekundách Použij `time.Now`
	// s `Unix` nebo `UnixNano`.
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// Všimni si že neexistuje nic jako `UnixMillis`, takže
	//  pro získání milisekund od epochy budeš muset ručně
	// vydělit nanosekundy.
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// Také můžeš zkonvertovat integer sekund nebo nanosekund
	// od epochy, to vrátí odpovídající `time`.
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
