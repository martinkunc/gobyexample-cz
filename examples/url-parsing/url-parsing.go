// URL poskytují [jednotný způsob pro adresování zdrojů](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// Tady je jak parsovat URL v Go.

package main

import "fmt"
import "net"
import "net/url"

func main() {

	// Budeme parsovat tento příklad URL, obsahující
	// schéma, informace o autentizaci, host, port, cestu,
	// parametry dotazu, a fragment dotazu.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parsuj URL a ujisti se že nenastaly žádné chyby.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Získání schématu je přímočaré.
	fmt.Println(u.Scheme)

	// `User` obsahuje všechny uživatelské informace o
	// autentizaci; Zavolej na něm `Username` a `Password`
	// pro konkrétní hodnoty.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host` obsahuje obojí, hostname a port,
	// pokud existují. Pro jejich exktrahování použij `SplitHostPort`.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Zde extrahujeme `path` a fragment za `#`.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// Pro získání parametrů dotazu v řetězci ve formátu  `k=v`,
	// použij `RawQuery`. Také můžeš rozparsovat parametry dotazu
	// do mapy. Mapa parsovaných parametrů dotazu jsou ze
	// řetězců na výřezy řetězců, takže indexuj do `[0]`
	// pokud chceš jenom první hodnotu.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
