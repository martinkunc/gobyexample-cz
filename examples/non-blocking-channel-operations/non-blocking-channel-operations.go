// Základní posílací a přijímací operace s kanály jsou
// blokující.
// Ačkoliv, můžeme použít `select` s klauzulí `default`
// pro implementací _neblokkujícího_ posílání, příjmů,
// neblokujících více-cestných `select`ů.

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Zde máme neblokující příjem.. Pokud je na
	// `messages` dostupná nějaká hodnota, tak s ní
	// `select` bude pokračovat `větví` `<-messages`.
	// A když ne, bude hned pokračovat větví
	// `default`.
	select {
	case msg := <-messages:
		fmt.Println("přijata zpráva", msg)
	default:
		fmt.Println("žádná zpráva nepřijata")
	}

	// Blokující posálíná funguje podobně. Níže, zpráva
	// `msg` nemůže být poslaná do kanálu `messages`,
	// protože ten nemá žádný buffer a není tu žádný
	// příjemce. Je tedy vybrána větev `default`.
	msg := "ahoj"
	select {
	case messages <- msg:
		fmt.Println("poslaná zpráva", msg)
	default:
		fmt.Println("žádná zpráva neposlaná")
	}

	// Také můžeme použít víc větví `case` nad klauzulí
	// `default` pro implementaci více-cestného
	// neblokujícího selektu. Zde zkoušíme neblokující
	// příjmy na obou, `messages` i `signals`.
	select {
	case msg := <-messages:
		fmt.Println("přijatá zpráva", msg)
	case sig := <-signals:
		fmt.Println("přijatý signál", sig)
	default:
		fmt.Println("žádná aktivita")
	}
}
