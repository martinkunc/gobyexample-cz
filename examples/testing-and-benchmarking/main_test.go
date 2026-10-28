// Unit testy jsou důležitou součástí psaní
// promyšlených Go programů. Balíček `testing`
// poskytuje nástroje, které pro psaní unit testů
// potřebujeme, a příkaz `go test` testy spouští.

// Pro účely ukázky je tento kód v balíčku `main`, ale
// mohl by být v libovolném balíčku. Testovací kód je
// obvykle ve stejném balíčku jako kód, který testuje.
package main

import (
	"fmt"
	"testing"
)

// Budeme testovat tuto jednoduchou implementaci minima
// dvou celých čísel. Kód, který testujeme, by obvykle
// byl ve zdrojovém souboru pojmenovaném třeba
// `intutils.go` a testovací soubor k němu by se pak
// jmenoval `intutils_test.go`.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Test vytvoříme napsáním funkce, jejíž název začíná
// na `Test`.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` ohlásí selhání testu, ale v provádění
		// testu pokračuje. `t.Fatal*` ohlásí selhání testu
		// a test okamžitě zastaví.
		t.Errorf("IntMin(2, -2) = %d; očekáváno -2", ans)
	}
}

// Psaní testů se může opakovat, proto je idiomatické
// používat *styl řízený tabulkou* (table-driven), kde
// jsou vstupy testů a očekávané výstupy uvedené
// v tabulce a jeden cyklus je projde a provede logiku
// testu.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// `t.Run` umožňuje spouštět "podtesty", jeden pro
		// každý záznam v tabulce. Ty se při spuštění
		// `go test -v` zobrazují samostatně.
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("máme %d, chceme %d", ans, tt.want)
			}
		})
	}
}

// Benchmarky jsou obvykle v souborech `_test.go` a jejich
// název začíná na `Benchmark`.
// Jakýkoli kód, který je pro běh benchmarku potřeba, ale
// neměl by se měřit, patří před tento cyklus.
func BenchmarkIntMin(b *testing.B) {
	for b.Loop() {
		// Spouštěč benchmarků automaticky provede tělo
		// tohoto cyklu mnohokrát, aby získal rozumný odhad
		// doby běhu jedné iterace.
		IntMin(1, 2)
	}
}
