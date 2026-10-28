// S použitím balíčku `net/http` je jednoduché
// napsat základní HTTP server.
package main

import (
	"fmt"
	"net/http"
)

// Základním konceptem v serverech v `net/http` jsou
// *handlery*. Handler je objekt implementující interface
// `http.Handler`. Běžný způsob pro psaní handleru
// je použít adapter `http.HandlerFunc` adaptér
// na funkcích s vhodnou signaturou.
func hello(w http.ResponseWriter, req *http.Request) {

	// Funkce sloužící jako handlery akceptují jako
	// parametry `http.ResponseWriter` a `http.Request`.
	// ResponseWriter je používaný pro plnění
	// HTTP odpovědi. Tady je naše jednoduchá odpověď
	// jenom "hello\n".
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// Tento handler dělá něco trochu sofistikovanějšího
	// když načte všechny hlavičky z požadavku a
	// vypisuje je do těla odpovědi.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// Registrujeme naše handlery k cestám na serveru s
	// použitím vhodné funkce `http.HandleFunc`. Nastavuje
	// *výchozí router* v balíčku `net/http` a bere funkci
	// jako argument.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Nakonec zavoláme `ListenAndServe` s portem a handlerem.
	// `nil` ji řekne aby použila výchozí router
	// který jsme právě nastavli.
	http.ListenAndServe(":8090", nil)
}
