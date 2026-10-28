# Spuštěné programy vrací výstup který je stejný 
# jako bychom je spustili přímo z příkazové řádky.
$ go run spawning-processes.go 
> date
Thu 05 May 2022 10:10:12 PM PDT

# date nemá volbu `-x`, takže skončí s chybovou
# hláškou a nenulovým návratovým kódem.
návratový kód příkazu = 1
> grep hello
hello grep

> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go
