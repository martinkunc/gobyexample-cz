# Spuštění tohoto programu ukazuje že první operace
# timeoutuje a druhá uspěje.
$ go run timeouts.go 
timeout 1
result 2

# Použití vzoru `select` timeoutu potřebuje
# poslání výsledků pomocí kanálů. Je to obecně dobrý
# nápad - další dúležité vlastnosti Go jsou na kanálech a
# `select`u založené. Příště se podíváme na dva takové
# příklady: časovače a tikače.
