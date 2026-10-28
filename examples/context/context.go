// V předchozím příkladu jsme se podívali na vytvoření
// jednoduchého [HTTP serveru](http-server). HTTP servery
// se hodí pro ukázku použití `context.Context` pro
// řízení rušení. `Context` přenáší termíny (deadlines),
// signály zrušení a další hodnoty vázané na požadavek
// přes hranice API a gorutin.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// `context.Context` vytváří pro každý požadavek
	// mechanismus `net/http` a je dostupný pomocí metody
	// `Context()`.
	ctx := req.Context()
	fmt.Println("server: hello handler nastartoval")
	defer fmt.Println("server: hello handler skončil")

	// Před odesláním odpovědi klientovi počkáme několik
	// sekund. To může simulovat nějakou práci, kterou
	// server dělá. Během práce sledujeme kanál `Done()`
	// kontextu, jestli nepřišel signál, že máme práci
	// zrušit a co nejdříve se vrátit.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// Metoda `Err()` kontextu vrací chybu, která
		// vysvětluje, proč byl kanál `Done()` zavřen.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// Stejně jako předtím zaregistrujeme náš handler na
	// cestu "/hello" a spustíme server.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
