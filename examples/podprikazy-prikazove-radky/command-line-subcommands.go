// Některé nástroje příkazové řádky, jako nástroj `go`
// nebo `git` mají mnoho *podpříkazů*, kažý se svou
// vlastní množinou voleb. Například, `go build` a
// `go get` jsou dva různé podpříkazy nástroje `go`.
// Balíček `flag` nám umožňuje snadno definovat
// jednoduché podpříkazy které mají své vlastní volby.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Podpříkaz deklarujeme s použitím funkce
	// `NewFlagSet`, a pokračujeme s definováním jednotlivých
	// voleb specifických pro tento podpříkaz.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// Pro jiné podpříkazy můžeme definovat jiné
	// podporované volby.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Podpříkaz je očekávaný jako první argument
	// programu.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Zjistíme jaký podpříkaz byl vyvolaný.
	switch os.Args[1] {

	// Pro každý podpříkaz naparsujeme jeho vlastní volby
	// a máme přístup ke zbývajícím pozičním argumentům.
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
