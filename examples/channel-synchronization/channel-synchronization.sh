$ go run channel-synchronization.go      
pracuji...hotovo                

# Když byste odstranili řádek `<- done`, program by
# mohl skončit dřív, než pracovní gorutina dokončí svou
# práci, a v některých případech dokonce ještě dřív, než
# by vůbec nastartovala.

