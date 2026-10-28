# Když spustíme tento program, nejdřív uvidíme výstup
# blokujícího volání, potom výstup ze dvou gorutin.
# Výstup gorutin může být prokládaný, protože gorutiny
# jsou spuštěné prováděcím prostředím Go souběžně.
$ go run goroutines.go
přímo : 0
přímo : 1
přímo : 2
gorutina : 0
gorutina : 1
gorutina : 2
jdeme
hotovo

# Jako další se podíváme na doplněk ke gorutinám v
# konkurentních programech: kanály.
