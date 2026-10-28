# Řetězec kóduje do lehce odlišných hodnot u standarního
# a URL base64 enkodóru (koncové `+` vs `-`)
# ale jak se očekává oba dekódují do původního řetězce.
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
