# Pro spuštění programu, vlož tento kód do
# `hello-world.go` a použij `go run`.
$ go run hello-world.go
ahoj světe

# Někdy budeme chtít vytvořit (zbuildovat) z našich
# programů binárky. To můžeme udělat s použitím 
# `go build`.
$ go build hello-world.go
$ ls
hello-world	hello-world.go

# Potom můžeme vytvořenou binárku přímo spustit.
$ ./hello-world
ahoj světe

# Teď když můžeme spouštět a kompilovat první Go programy,
# pojďme se naučit víc o jazyku samotném.
