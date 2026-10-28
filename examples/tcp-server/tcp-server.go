// Balíček `net` poskytuje nástroje, které potřebujeme
// pro snadné vytvoření serverů s TCP sockety.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	// `net.Listen` spustí server na dané síti (TCP)
	// a adrese (port 8090 na všech rozhraních).
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Chyba při naslouchání:", err)
	}

	// Při ukončení aplikace listener zavřeme, aby se
	// port uvolnil.
	defer listener.Close()

	// Nekonečný cyklus pro přijímání nových klientských
	// spojení.
	for {
		// Čekáme na spojení.
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Chyba při přijímání spojení:", err)
			continue
		}

		// Pro obsluhu spojení zde použijeme gorutinu, aby
		// hlavní cyklus mohl dál přijímat další spojení.
		go handleConnection(conn)
	}
}

// `handleConnection` obsluhuje jedno klientské spojení:
// načte od klienta jeden řádek textu a vrátí odpověď.
func handleConnection(conn net.Conn) {
	// Zavřením spojení uvolníme prostředky, když
	// komunikaci s klientem dokončíme.
	defer conn.Close()

	// Pomocí `bufio.NewReader` načteme od klienta jeden
	// řádek dat (ukončený znakem nového řádku).
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Chyba čtení: %v", err)
		return
	}

	// Vytvoříme a pošleme klientovi zpět odpověď, čímž
	// ukážeme obousměrnou komunikaci.
	ackMsg := strings.ToUpper(strings.TrimSpace(message))
	response := fmt.Sprintf("ACK: %s\n", ackMsg)
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Chyba zápisu na serveru: %v", err)
	}
}
