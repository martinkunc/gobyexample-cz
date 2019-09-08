# Po spuštětní našeho programu uvidíme první dávku
# požadavků zpracovanou jednou za ~200 milisekund,
# jak jsme chtěli.
$ go run rate-limiting.go
požadavek 1 2009-11-10 23:00:00.2 +0000 UTC m=+0.200000001
požadavek 2 2009-11-10 23:00:00.4 +0000 UTC m=+0.400000001
požadavek 3 2009-11-10 23:00:00.6 +0000 UTC m=+0.600000001
požadavek 4 2009-11-10 23:00:00.8 +0000 UTC m=+0.800000001
požadavek 5 2009-11-10 23:00:01 +0000 UTC m=+1.000000001

# Ve druhé dávce požadavků obsloužíme první
# 3 okamžitě díky velikosti dávky v limitování rychlosti,
# potom obsloužíme zbývající 2 s přibližnou prodlevou
# ~200ms na každý.
požadavek 1 2009-11-10 23:00:01 +0000 UTC m=+1.000000001
požadavek 2 2009-11-10 23:00:01 +0000 UTC m=+1.000000001
požadavek 3 2009-11-10 23:00:01 +0000 UTC m=+1.000000001
požadavek 4 2009-11-10 23:00:01.2 +0000 UTC m=+1.200000001
požadavek 5 2009-11-10 23:00:01.4 +0000 UTC m=+1.400000001
