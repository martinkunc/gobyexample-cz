# Spuštění programu vypočítá hash a vypíše ho ve
# lidsky čitelném hexa formátu.
$ go run sha256-hashes.go
sha256 this string
1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a...


# S použitím podobného vzoru můžete vypočítat jiné hashe
# podobné tomu výše. Například pro výpočet SHA512 hashe
# importujte `crypto/sha512` a použijte `sha512.New()`.

# Poznamenejte si, že pokud budete potřebovat kryptograficky
# bezpečné hashe, měli byste opatrně prozkoumat
# [sílu hash funkcí](https://en.wikipedia.org/wiki/Cryptographic_hash_function)!
