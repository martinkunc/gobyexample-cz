# Spustí server na pozadí.
$ go run context.go &

# Nasimulujeme požadavek klienta na `/hello` a krátce po
# spuštění stiskneme Ctrl+C, čímž signalizujeme zrušení.
$ curl localhost:8090/hello
server: hello handler nastartoval
^C
server: context canceled
server: hello handler skončil
