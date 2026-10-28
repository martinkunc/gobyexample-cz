# Všimněte si, že i když jsou výřezy jiné typy než pole,
# jsou pomocí `fmt.Println` renderované podobně.
$ go run slices.go
prázdné: [  ]
nastavené: [a b c]
získáno: c
délka: 3
přidané: [a b c d e f]
kopie: [a b c d e f]
vy1: [c d e]
vy2: [a b c d e]
vy3: [c d e f]
dekl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

# Pro další detaily o návrhu a implementaci výřezů v Go se
# podívejte na tento [skvělý blogpost](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html)
# od Go týmu.

# Teď když jsme viděli pole a výřezy, podíváme se na
# další klíčovou zabudovanou datovou strukturu v Go:
# mapy.
