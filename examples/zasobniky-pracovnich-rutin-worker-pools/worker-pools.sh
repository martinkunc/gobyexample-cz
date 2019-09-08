# Náš běžící program ukazuje že se provádí 5 úkolů různými
# pracovními rutinami. Program běžá jenom asi 2 vteřiny
# ačkoliv pracujeme na celkové práci okolo 5 vteřin, to je
# proto, že 3 rutiny běží souběžné.
$ time go run worker-pools.go 
worker 3 nastartoval job 1
worker 1 nastartoval job 2
worker 2 nastartoval job 3
worker 1 ukončil job 2
worker 1 nastartoval job 4
worker 3 ukončil job 1
worker 3 nastartoval job 5
worker 2 ukončil job 3
worker 3 ukončil job 5
worker 1 ukončil job 4

real	0m2.358s
