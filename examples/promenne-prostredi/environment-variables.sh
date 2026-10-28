# Spuštění tohoto programu ukazuje že jsme načetli
# hodnotu pro `FOO` kterou jsme nastavili v programu, ale
# že `BAR` je prázdné.
$ go run environment-variables.go
FOO: 1
BAR: 

# Seznam klíčů v prostředí bude záležet na
# konkrétním stroji.
TERM_PROGRAM
PATH
SHELL
...

# Když nejdřív nastavíme v prostředí `BAR`, spuštění
# programu poté hodnotu načte.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
