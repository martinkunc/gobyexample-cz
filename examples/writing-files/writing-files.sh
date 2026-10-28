# Zkus spuštění kódu pro zápis do souboru.
$ go run writing-files.go 
zapsáno 5 bajtů
zapsáno 7 bajtů
zapsáno 13 bajtů

# Potom zkontroluj obsah zapsaných souborů.
$ cat /tmp/dat1
ahoj
go
$ cat /tmp/dat2
nějaké
zápisy
bufferovány

# Příště se podíváme na některé vlastností I/O které jsme
# právě viděli u streamů `stdin` a `stdout`.
