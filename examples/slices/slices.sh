# Všimněte si, že i když jsou výřezy jiné typy než pole,
# jsou pomocí `fmt.Println` renderované podobně.
$ go run slices.go
neinic: [] true true
prázdné: [  ] délka: 3 kap.: 3
nastavené: [a b c]
získáno: c
délka: 3
přidané: [a b c d e f]
kopie: [a b c d e f]
vy1: [c d e]
vy2: [a b c d e]
vy3: [c d e f]
dekl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]

# Pro další detaily o návrhu a implementaci výřezů v Go se
# podívejte na tento [skvělý blogpost](https://go.dev/blog/slices-intro)
# od Go týmu.

# Teď když jsme viděli pole a výřezy, podíváme se na
# další klíčovou zabudovanou datovou strukturu v Go:
# mapy.
