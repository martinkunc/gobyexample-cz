# Spuštění tohoto programu způsobí že spadne na paniku,
# vypíše chybu a stacktrace gorutiny, a skončí s
# ne-nulovým stavem.
$ go run panic.go
panic: problém

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# Všimněte si, že zatímco v některých jazycích které
# používají výjimky pro ošetřování mnoha chyb, v Go je
# idiomatické použití návratových hodnot vracejících chybu
# všude kde je to možné.
