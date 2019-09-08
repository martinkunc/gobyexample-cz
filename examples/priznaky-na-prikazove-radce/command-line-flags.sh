# Pro experimentování s volbami na příkazové řádce je
# nejlepší program nejdřív zkompilovat a pak spouštět
# výslednou binárku přímo.
$ go build command-line-flags.go

# Vyzkoušíme zbuildovaný program nejdřív poskytnutím 
# hodnot pro všechny volby.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Všimněte si, že když volby vynecháme, dostanou 
# automaticky své výchozí hodnoty.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# Zbývající poziční argumenty je možné poskytnout po
# nějakých volbách.
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# Všimněte si, že balíček `flag` vyžaduje aby všechny 
# se všechny volby nastavily před pozičními argumenty
# (jinak budou volby interpretované jako poziční 
# argumenty).
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# Použij volby `-h` nebo `--help` pro získání automaticky
# generovaného textu nápovědy pro přkazovou řádku
# programu.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Pokud použiješ volbu která nebyla předem specifikovaná
# balíčku `flag`, program vypíše chybovou hlášku a
# ukáže znovu text nápovědy.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
