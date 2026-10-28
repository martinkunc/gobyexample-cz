#  Když spustíte `exit.go` s použitím `go run`, návratová
# hodnota bude samotným `go` vypsána.
$ go run exit.go
exit status 3

# Když zbuildujete a spustíte binárku, můžete si zjistit
# návratový kód z terminálu
$ go build exit.go
$ ./exit
$ echo $?
3

# Všimněte si, že `!` z našeho programu není nikde
# vypsaný.
