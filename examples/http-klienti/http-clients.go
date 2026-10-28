// Standardní klihovna Go přichází s excelentní podporou
// pro HTTP klienty a servery v balíčku `net/http`.
// V tomto příkladu ho použijeme pro vytvoření
// jednoduchých HTTP požadavků.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Vytvoř HTTP GET požadavek na server. `http.Get` je
	// vhodnou zkratkou okolo vytváření objektu `http.Client`
	// a volání jeho metody `Get`; používá objekt
	// `http.DefaultClient` object který má užitečná
	// výchozí nastavení.
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Vypíše status HTTP odpovědi.
	fmt.Println("Response status:", resp.Status)

	// Vypíše prvních 5 řádků těla odpovědi.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
