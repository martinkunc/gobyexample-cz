# Když náš program zapneme zpráva `"ping"` je úspěšně
# předaná z jedné gorutiny do jiné pomocí našeho kanálu.

$ go run channels.go 
ping

# Běžné poslání a přijímání blokuje, dokud oba
# odesílající i příjemce nejsou připravení. Toto chování
# způsobila že jsme čekali na konci našeho programu na
# zprávu`"ping"` bez nutnosti použít nějakou jinou
# synchronizaci.
