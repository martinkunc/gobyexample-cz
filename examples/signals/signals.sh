# Když spustíme tento program, bude blokovat a čekat
# na signál. Stisknutím `ctrl-C` (které se na terminálu
# ukáže jako `^C`) můžeme poslat signál `SIGINT`,
# který způsobí, že program vypíše příčinu zrušení a poté
# se ukončí.
$ go run signals.go
awaiting signal
^C
interrupt signal received
exiting
