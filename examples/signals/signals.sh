# Když spustíme tento program tak bude blokovat a čekat 
# na signál. Stisknutím `ctrl-C` (které se na terminálu
# ukáže jako `^C`) můžeme poslat signál `SIGINT`,
# který způsobí že program vypíše `interrupt` a poté se
# ukončí.
$ go run signals.go
awaiting signal
^C
interrupt
exiting
