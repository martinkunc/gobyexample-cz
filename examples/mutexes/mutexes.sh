# Spuštění programu ukáže, že se countery
# aktualizovaly podle očekávání.
$ go run mutexes.go
map[a:20000 b:10000]

# Dále se podíváme na implementaci tohoto stejného úkolu
# se správou stavu jenom s pomocí gorutin a kanálů.
