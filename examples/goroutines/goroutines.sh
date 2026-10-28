# Když zapneme tento program, nejdřív uvidíme výstup
# blokujícího volání, potom uvidíme prokládaný výstup ze
# dvou gorutin. Toto prokládání odráží že gorutiny jsou
# spuštěné prováděcím prostředím Go současně.
$ go run goroutines.go
přímo : 0
přímo : 1
přímo : 2
gorutina : 0
gorutina : 1
gorutina : 2
jdeme
<enter>
hotovo


# Jako další se podíváme na doplněk ke gorutinám v
# konkurentních programech: kanály.
