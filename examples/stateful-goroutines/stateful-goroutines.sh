# Spuštění našeho programu ukazuje že správa stavu
# založená na gorutině provedla celkem přibližně 80,000
# operací.
$ go run stateful-goroutines.go
readOps: 71708
writeOps: 7177

# V tomto konkrétním případě byl přístup založený na
# gorutinách trochu komplikovanější než přístup založený
# na mutexu. Přesto ale může být v některých případech
# užitečný, například když jsou zapojeny i další kanály,
# nebo když by správa více takových mutexů byla náchylná
# k chybám. Měl/a bys zvolit takový přístup který se zdá
# přirozenější, obzvlášť s respektem k pochopení 
# správnosti tvého programu.
