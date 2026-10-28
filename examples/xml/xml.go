// Go nabízí zabudovanou podporu pro XML a formáty
// podobné XML pomocí balíčku `encoding/xml`.

package main

import (
	"encoding/xml"
	"fmt"
)

// Struktura Plant bude namapována na XML. Podobně jako
// v příkladech s JSON obsahují značky (tagy) prvků
// direktivy pro enkodér a dekodér. Tady používáme
// některé speciální vlastnosti balíčku XML: název prvku
// `XMLName` určuje název XML elementu, který tuto
// strukturu reprezentuje; `id,attr` znamená, že prvek
// `Id` je XML _atribut_, a ne vnořený element.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Rostlina id=%v, název=%v, původ=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Káva"}
	coffee.Origin = []string{"Etiopie", "Brazílie"}

	// Vygenerujeme XML reprezentující naši rostlinu;
	// pomocí `MarshalIndent` dostaneme výstup lépe
	// čitelný pro člověka.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// Pro přidání obecné XML hlavičky do výstupu ji
	// explicitně připojte.
	fmt.Println(xml.Header + string(out))

	// Pomocí `Unmarshal` naparsujeme proud bajtů s XML
	// do datové struktury. Pokud je XML poškozené nebo
	// ho nelze namapovat na Plant, vrátí se popisná chyba.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Rajče"}
	tomato.Origin = []string{"Mexiko", "Kalifornie"}

	// Značka prvku `parent>child>plant` říká enkodéru,
	// aby všechny `plant` vnořil pod `<parent><child>...`
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
