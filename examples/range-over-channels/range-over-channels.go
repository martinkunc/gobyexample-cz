// V [předchozím](range) příkladu jsme viděli jak `for` a
// `range` umožňují iterování přes základní datové
// struktury. Tento zápis můžeme také použít pro iterování
// po hodnotách přijímaných z kanálu.

package main

import "fmt"

func main() {

	// Budeme iterovat přes 2 hodnoty v kanálu `queue`.
	queue := make(chan string, 2)
	queue <- "jedna"
	queue <- "dvě"
	close(queue)

	// Tento `range` iteruje přes každý element jak jsou
	// přijímané z kanálu `queue`. Protože jsme kanál
	// výše zavřeli s pomocí `close`, iterace po přijetí
	// dvou prvků skončí.
	for elem := range queue {
		fmt.Println(elem)
	}
}
