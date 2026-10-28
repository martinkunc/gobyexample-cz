// Go nabízí zabudovanou podporu pro JSON kódování
// a dekódování, včetně do a ze zabudovaných i
// uživatelských datových typů.

package main

import (
	"bytes"
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Tyto dvě struktury použijeme k demonstrování kódování a
// dekódování uživatelských datových typů níže.
type response1 struct {
	Page   int
	Fruits []string
}

// Do JSON se kódují/dekódují jen exportované prvky.
// Aby byly prvky exportované, musí začínat velkým písmenem.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// Nejprve se podíváme na kódování základních datových typů
	// do řetězců JSON. Zde jsou nějaké příklady pro atomické
	// hodnoty.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// A zde jsou nějaké příklady pro výřezy a mapy, které se
	// kódují do JSON polí a objektů, jak byste asi čekali.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Balíček JSON umí automaticky kódovat i vaše
	// uživatelské datové typy. Do kódovaného výstupu
	// se zahrnou jen exportované prvky a ve výchozím
	// stavu se použijí jejich názvy jako JSON klíče.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Na deklaraci prvků struktur můžete použít značky (tagy)
	// pro přizpůsobení názvů zakódovaných JSON klíčů. Podívejte
	// se na definici struktury `response2` výše pro příklad
	// takových značek.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Teď se pojďme podívat na dekódování JSON dat do Go
	// hodnot. Zde je příklad pro obecnou datovou
	// strukturu.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Musíme poskytnout nějakou proměnnou kam bude
	// balíček JSON vkládat dekódovaná data. Typ
	// `map[string]any` bude obsahovat mapu řetězců
	// na libovolný datový typ.
	var dat map[string]any

	// Tady je vlastní dekódování a kontrola s tím spojených
	// chyb.
	// Pro stručnost v těchto příkladech chyby ignorujeme;
	// ve skutečném kódu byste ale chyby měli vždy
	// kontrolovat a reagovat na ně.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Abychom mohli hodnoty v dekódované mapě použít,
	// budeme je muset zkonvertovat do vhodného typu.
	// Například zde konvertujeme hodnotu v `num` do
	// očekávaného typu `float64`.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Přístup ke vnořeným datům vyžaduje sérii
	// konverzí.
	strs := dat["strs"].([]any)
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Také můžeme dekódovat JSON do uživatelských datových typů.
	// Výhoda je větší typová bezpečnost v našich programech a
	// odstranění potřeby typových vynucení při přístupu k dekódovaným
	// datům.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	_ = json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// V příkladech výše jsme vždy používali bajty a řetězce
	// jako meziprodukty mezi daty a JSON reprezentací na
	// standardním výstupu. Také můžeme streamovat JSON
	// kódování přímo do `io.Writer`u. Například do
	// `os.Stdout` a nebo dokonce do těla HTTP odpovědi.
	d := map[string]int{"apple": 5, "lettuce": 7}
	var buf bytes.Buffer
	_ = json.MarshalWrite(&buf, d)
	fmt.Println(buf.String())

	// Streamované čtení z `io.Reader`ů, jako je `os.Stdin`
	// nebo tělo HTTP požadavku, se provádí pomocí
	// `json.UnmarshalRead`.
	res1 := response2{}
	_ = json.UnmarshalRead(strings.NewReader(str), &res1)
	fmt.Println(res1)
}
