# Očekáváme přesně 50 000 operací. Kdybychom použili
# neatomický integer a zvyšovali ho pomocí `ops++`,
# pravděpodobně bychom dostali jiné číslo, měnící se
# mezi jednotlivými spuštěními, protože by se gorutiny
# navzájem ovlivňovaly. Navíc bychom při spuštění
# s příznakem `-race` dostali chyby souběhu (data race).
$ go run atomic-counters.go
ops: 50000

# Jako další se podíváme na mutexy, které jsou dalším
# nástrojem pro správu stavu.
