# Spuštění programu ukáže že jsme provedli celkem okolo
# 90,000 operací proti našemu `stavu` synchronizovaného
# `mutexem`.
$ go run mutexes.go
readOps: 83285
writeOps: 8320
stav: map[1:97 4:53 0:33 2:15 3:2]

# Dále se podíváme na implementaci tohoto stejného úkolu 
# se správou stavu jenom s pomocí gorutin a kanálů.
