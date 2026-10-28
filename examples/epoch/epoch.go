// Běžný požadavek v programech je získání počtu vteřin,
// milisekund, nebo nanosekund z
// [Unixového času](https://en.wikipedia.org/wiki/Unix_time).
// Tady je, jak to udělat v Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Pro získání uplynulého času od Unix epochy v
	// sekundách, milisekundách nebo nanosekundách použijte
	// `time.Now` s `Unix`, `UnixMilli` nebo `UnixNano`.
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// Také můžete zkonvertovat integer sekund nebo nanosekund
	// od epochy na odpovídající `time`.
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
