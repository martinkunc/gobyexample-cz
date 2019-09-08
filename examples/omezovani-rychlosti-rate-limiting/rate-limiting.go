// <em>[Omezování rychlosti (Rate limiting)](http://cs.wikipedia.org/wiki/Rate_limiting)</em>
// je důležitý mechanismus pro kontrolu využívání zdrojů
// a udržování kvality služeb. Go elegantně podporuje
// omezování rychlostí pomocí gorutin,
// kanálů a [opakujících časovaců tickers](opakujici-casovace-tickers).

package main

import "time"
import "fmt"

func main() {

	// Nejprve se podíváme na základní omezení rychlosti.
	// Předpokládejme že chceme omezit správu našich
	// příchozích požadavků.
	// Budeme obsluhovat tyto požadavky z kanálu se jménem
	// requests.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Tento `limiter` kanál přijme hodnotu
	// kažých 200 milisekund. Je to regulátor v našem
	// schématu pro omezovámí rychlosti.
	limiter := time.Tick(200 * time.Millisecond)

	// Pomocí blokování na příjmu z kanálu `limiter`
	// před obsluhou každého pořadavku, se omezíme na
	// 1 požadavek každých 200 milisekund.
	for req := range requests {
		<-limiter
		fmt.Println("požadavek", req, time.Now())
	}

	// Možná že chceme dovolit krátké dávky pořadavků v
	// našem schématu omezování rychlosti. zatímco
	// zachováváme celkový limit rychlosti.
	// To můžeme udlělat pomocí kanálu s vyrovnávací
	// pamětí. Tento kanál `burstyLimiter`
	// dovolí dávku až do 3 událostí.
	burstyLimiter := make(chan time.Time, 3)

	// Naplníme kanál pro demonstraci dovolené velikosti
	// dávky.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Kažých 200 milisekund se pokusíme přidat novou
	// hodnotu do kanálu `burstyLimiter`, až do jeho
	// limitu 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Nyní simulujeme 5 dalších příchozích požadavků.
	// První tři z nich využijí dávkovou schopnost
	// `burstyLimiter`u.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("požadavek", req, time.Now())
	}
}
