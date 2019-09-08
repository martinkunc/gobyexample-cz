# Spuštění našeho programu ukáže list setříyení podle
# délky řetězce, tak jak jsme chtěli.
$ go run sorting-by-functions.go 
[kiwi peach banana]

# Na tomto principu, vytvořením uživatelského typu,
# implementováním tří metod `Interface` na tomto typu a
# poté zavoláním funkce sort.Sort na kolekci tohoto
# uživatelského typu, můžeme setřídit výřezy v Go podle
# libovolných funkcí.
