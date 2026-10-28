# Když spustíme náš program, je nahrazený `ls`.
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Všimněte si, že Go nenabízí klasickou Unix funkci
# `fork`. Obvykle to ale není problém, protože spouštění 
# gorutin, vztváření podprocesů, a spouštění procesů
# pokrývá většinu případů použití `fork`u.
