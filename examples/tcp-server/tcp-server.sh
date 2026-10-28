# Spustí TCP server na pozadí.
$ go run tcp-server.go &

# Pošle data a zachytí odpověď pomocí netcat.
$ echo "Ahoj z netcat" | nc localhost 8090
ACK: AHOJ Z NETCAT
