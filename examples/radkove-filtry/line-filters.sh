# Vyzkoušej náš řádkový filtry, nejdřív udělej soubor
# s pár malými znaky.
$ echo 'ahoj'   > /tmp/lines
$ echo 'filtre' >> /tmp/lines

# Potom použij řádkový filtr pro získání kapitalizovaných
# řádků.
$ cat /tmp/lines | go run line-filters.go
AHOJ
FILTRE
