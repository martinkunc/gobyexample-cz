# Zkus spuštění kódu pro zápis do souboru.
$ go run writing-files.go 
zapsáno 9 bajtů
zapsáno 7 bajtů
zapsáno 13 bajtů

# Potom zkontroluj obsah zapsaných souborů.
$ cat /tmp/dat1
ahoj
go
$ cat /tmp/dat2
nějaké
writes
bufferováno

# Příště se podíváme na některé vlastností I/O které jsme
# právě viděli u streamů `stdin` a `stdout`.
