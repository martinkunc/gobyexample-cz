# Ukázkový výstup; vypsané datum a čas budou záviset na
# tom, kdy byl příklad spuštěn.
$ go run logging.go
2023/08/22 10:45:16 standardní logger
2023/08/22 10:45:16.904141 s mikrosekundami
2023/08/22 10:45:16 logging.go:39: se souborem/řádkem
my:2023/08/22 10:45:16 z mylog
ohmy:2023/08/22 10:45:16 z mylog
z buflog:buf:2023/08/22 10:45:16 ahoj

# Tyto řádky jsou pro přehlednost na webu zalomené;
# ve skutečnosti se vypisují na jednom řádku.
{"time":"2023-08-22T10:45:16.904166391-07:00",
 "level":"INFO","msg":"ahoj"}
{"time":"2023-08-22T10:45:16.904178985-07:00",
	"level":"INFO","msg":"znovu ahoj",
	"key":"val","age":25}
