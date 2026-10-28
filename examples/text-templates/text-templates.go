// Go nabízí zabudovanou podporu pro vytváření dynamického
// obsahu nebo zobrazování přizpůsobeného výstupu
// uživateli pomocí balíčku `text/template`. Sesterský
// balíček `html/template` poskytuje stejné API, ale má
// navíc bezpečnostní funkce a měl by se používat pro
// generování HTML.

package main

import (
	"os"
	"text/template"
)

func main() {

	// Můžeme vytvořit novou šablonu a naparsovat její tělo
	// z řetězce.
	// Šablony jsou směsí statického textu a "akcí"
	// uzavřených v `{{...}}`, které slouží k dynamickému
	// vkládání obsahu.
	t1 := template.New("t1")
	t1, err := t1.Parse("Hodnota je {{.}}\n")
	if err != nil {
		panic(err)
	}

	// Případně můžeme použít funkci `template.Must`, která
	// vyvolá paniku, pokud `Parse` vrátí chybu. To se hodí
	// hlavně pro šablony inicializované v globálním
	// rozsahu.
	t1 = template.Must(t1.Parse("Hodnota: {{.}}\n"))

	// "Vykonáním" šablony vygenerujeme její text
	// s konkrétními hodnotami pro její akce. Akce `{{.}}`
	// je nahrazena hodnotou předanou jako parametr
	// do `Execute`.
	t1.Execute(os.Stdout, "nějaký text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// Pomocná funkce, kterou použijeme níže.
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// Pokud jsou data struktura, můžeme pro přístup k jejím
	// položkám použít akci `{{.FieldName}}`. Položky by
	// měly být exportované, aby byly při vykonávání
	// šablony dostupné.
	t2 := Create("t2", "Jméno: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jan Novák"})

	// Totéž platí pro mapy; u map nejsou žádná omezení na
	// velikost písmen v názvech klíčů.
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// if/else umožňuje v šablonách podmíněné vykonávání.
	// Hodnota se považuje za nepravdivou, pokud je to
	// výchozí hodnota typu, jako 0, prázdný řetězec,
	// nil ukazatel, atd.
	// Tento příklad ukazuje další vlastnost šablon:
	// použití `-` v akcích pro oříznutí bílých znaků.
	t3 := Create("t3",
		"{{if . -}} ano {{else -}} ne {{end}}\n")
	t3.Execute(os.Stdout, "neprázdný")
	t3.Execute(os.Stdout, "")

	// Bloky range nám umožňují procházet výřezy, pole, mapy
	// nebo kanály. Uvnitř bloku range je `{{.}}` nastaveno
	// na aktuální prvek iterace.
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
